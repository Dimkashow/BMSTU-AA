package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const TH_SIZE = 6
const DEBUG_INFO = false


var Md5Defender sync.Mutex

func CalcMd5(data string) string {
	Md5Defender.Lock()
	result := DataSignerMd5(data)
	Md5Defender.Unlock()
	return result
}

func CalcSingleHash(out chan string, str string) {
	out <- DataSignerCrc32(str)
}

func AsyncSingleHash(out chan interface{}, str, strMd5 string, SHash *sync.WaitGroup) {
	defer SHash.Done()

	ChanCrc32, ChanCrc32Md5 := make(chan string, 1), make(chan string, 1)
	go CalcSingleHash(ChanCrc32, str)
	go CalcSingleHash(ChanCrc32Md5, strMd5)

	strCrc32, strCrc32Md5 := <-ChanCrc32, <-ChanCrc32Md5
	out <- strCrc32 + "~" + strCrc32Md5
	if DEBUG_INFO {
		fmt.Println(str, " SingleHash data ", str)
		fmt.Println(str, " SingleHash md5(data) ", strMd5)
		fmt.Println(str, " SingleHash crc32(md5(data)) ", strCrc32Md5)
		fmt.Println(str, " SingleHash strCrc32 ", strCrc32)
		fmt.Println(str, " SingleHash result ", strCrc32+"~"+strCrc32Md5)
	}

}

// SingleHash считает значение crc32(data)+"~"+crc32(md5(data)) ( конкатенация двух строк через ~),
//  где data - то что пришло на вход (по сути - числа из первой функции)
func SingleHash(in, out chan interface{}) {
	SHash := &sync.WaitGroup{}
	for data := range in {
		dataToStr := fmt.Sprintf("%v", data)
		strMd5 := CalcMd5(dataToStr)
		SHash.Add(1)
		go AsyncSingleHash(out, dataToStr, strMd5, SHash)
	}
	SHash.Wait()
}

func CalcMultiHash(Arr *[]string, idx int, str string, wg *sync.WaitGroup) {
	defer wg.Done()
	(*Arr)[idx] = DataSignerCrc32(strconv.Itoa(idx) + str)
}

func AsyncMultiHash(out chan interface{}, data interface{}, MHash *sync.WaitGroup) {
	defer MHash.Done()

	str := fmt.Sprintf("%v", data)
	MHashInside := &sync.WaitGroup{}
	ArrFinalResult := make([]string, TH_SIZE)
	for TH := 0; TH < TH_SIZE; TH++ {
		MHashInside.Add(1)
		go CalcMultiHash(&ArrFinalResult, TH, str, MHashInside)
	}
	MHashInside.Wait()
	out <- strings.Join(ArrFinalResult, "")
	if DEBUG_INFO {
		for TH := 0; TH < TH_SIZE; TH++ {
			fmt.Println(str, " MultiHash result: crc32(th+step1) ", TH, ArrFinalResult[TH])
		}
		fmt.Println(str, " MultiHash result:", strings.Join(ArrFinalResult, ""))
	}
}

// * MultiHash считает значение crc32(th+data)) (конкатенация цифры, приведённой к строке и строки),
// где th=0..5 ( т.е. 6 хешей на каждое входящее значение ), потом берёт конкатенацию результатов в
//порядке расчета (0..5), где data - то что пришло на вход (и ушло на выход из SingleHash)
func MultiHash(in, out chan interface{}) {
	MHash := &sync.WaitGroup{}
	for data := range in {
		MHash.Add(1)
		go AsyncMultiHash(out, data, MHash)
	}
	MHash.Wait()
}

// * CombineResults получает все результаты, сортирует (https://golang.org/pkg/sort/),
// объединяет отсортированный результат через _ (символ подчеркивания) в одну строку
func CombineResults(in, out chan interface{}) {
	var result []string
	for el := range in {
		result = append(result, el.(string))
	}
	sort.Strings(result)
	strResult := strings.Join(result, "_")
	out <- strResult
	if DEBUG_INFO {
		fmt.Println("CombineResults", strResult)
	}
}

// * ExecutePipeline которая обеспечивает нам конвейерную обработку функций-воркеров, которые что-то делают.
func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	in := make(chan interface{})
	close(in)
	for _, jobF := range jobs {
		wg.Add(1)
		out := make(chan interface{})
		go func(in, out chan interface{}, wg *sync.WaitGroup, jobFunc job) {
			defer wg.Done()
			defer close(out)
			jobFunc(in, out)
		}(in, out, wg, jobF)
		in = out
	}
	wg.Wait()
}
