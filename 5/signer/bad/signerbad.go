package bad
/*
import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const TH_SIZE = 6
const DEBUG_INFO = false

// SingleHash считает значение crc32(data)+"~"+crc32(md5(data)) ( конкатенация двух строк через ~),
//  где data - то что пришло на вход (по сути - числа из первой функции)
func SingleHash(in, out chan interface{}) {
	for data := range in {
		dataToStr := fmt.Sprintf("%v", data)
		strMd5 := DataSignerMd5(dataToStr)
		strCrc32, strCrc32Md5 := DataSignerCrc32(dataToStr), DataSignerCrc32(strMd5)
		out <- strCrc32 + "~" + strCrc32Md5
	}
}

// * MultiHash считает значение crc32(th+data)) (конкатенация цифры, приведённой к строке и строки),
// где th=0..5 ( т.е. 6 хешей на каждое входящее значение ), потом берёт конкатенацию результатов в
//порядке расчета (0..5), где data - то что пришло на вход (и ушло на выход из SingleHash)
func MultiHash(in, out chan interface{}) {
	for data := range in {
		str := fmt.Sprintf("%v", data)
		ArrFinalResult := make([]string, TH_SIZE)
		for TH := 0; TH < TH_SIZE; TH++ {
			(ArrFinalResult)[TH] = DataSignerCrc32(strconv.Itoa(TH) + str)
		}
		out <- strings.Join(ArrFinalResult, "")
	}
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
*/