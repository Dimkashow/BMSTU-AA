#include <iostream>
#include <ctime>
#include <vector>

using namespace std;

int main()
{
    clock_t start = clock();

    const int AIM = 1000000;

    int startSize = AIM; // стартовый размер массива натуральных чисел
    int addSize = AIM; // размер дополнительного массива натуральных чисел

    vector<bool> nums(startSize);
    vector<int> primeNums(AIM);

    int foundPrimes = 0;

    for (int i = 2; i < startSize; i++)
        nums[i] = true;

    bool addition = false;
    int adder = 0;
    while (true)
    {
        if (addition)
        {
            nums.resize(nums.size() + addSize, true);

// исключим составные числа простыми из предыдущих итераций
            for (int i = 0; i < foundPrimes; i++)
            {
                int cur_num = primeNums[i];
                if ((addSize + ((nums.size() - addSize) % cur_num)) < cur_num)
                    continue;
                for (int j = ((nums.size() - addSize) / cur_num) * cur_num; j < nums.size(); j += cur_num)
                    nums[j] = false;
            }
        }
        else
            addition = true;


        int iter;
        if (foundPrimes == 0)
            iter = 2;
        else
            iter = primeNums[foundPrimes - 1] + 2;

        for ( ; iter < nums.size(); iter++)
        {
// выбираем очередное простое число
            if (nums[iter])
            {
                primeNums[foundPrimes] = iter;
                foundPrimes++;
                if (foundPrimes == AIM)
                    break;
// просеиваем
                for (int j = iter + iter; j < nums.size(); j += iter)
                    nums[j] = false;
            }
            else
                continue;

        }
        if (foundPrimes == AIM)
            break;
    }

    cout << "Last prime: " << primeNums[AIM -1];

    clock_t end = clock();
    cout << "\n----------------" << endl;
    cout << "Time: " << double(end - start) / CLOCKS_PER_SEC << " sec" << endl;
    return 0;
}