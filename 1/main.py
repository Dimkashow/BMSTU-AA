import string
import random
import time

OUTPUT_MODE = False


def RandomString(strLength=5):
    letters = string.ascii_lowercase
    return ''.join(random.choice(letters) for _ in range(strLength))


def OutputTable(table, str1, str2):
    print("\n   ", end=" ")
    for i in str2:
        print(i, end=" ")

    for i in range(len(table)):
        if i:
            print("\n" + str1[i - 1], end=" ")
        else:
            print("\n ", end=" ")
        for j in range(len(table[i])):
            print(table[i][j], end=" ")
    print("\n")


def LevRecursion(str1, str2, Dam=False):
    if str1 == '' or str2 == '':
        return abs(len(str1) - len(str2))
    forfeit = 0 if (str1[-1] == str2[-1]) else 1
    res = min(LevRecursion(str1, str2[:-1], Dam) + 1,
              LevRecursion(str1[:-1], str2, Dam) + 1,
              LevRecursion(str1[:-1], str2[:-1], Dam) + forfeit)
    if Dam:
        if len(str1) >= 2 and len(str2) >= 2 and str1[-1] == str2[-2] and str1[-2] == str2[-1]:
            res = min(res, LevRecursion(str1[:-2], str2[:-2], Dam) + 1)
    return res


def LevRecursionMatrix(str1, str2, matrix):
    if matrix[len(str1)][len(str2)] is not None:
        return matrix[len(str1)][len(str2)]

    if len(str1) == 0 or len(str2) == 0:
        matrix[len(str1)][len(str2)] = len(str2) + len(str1)
        return matrix[len(str1)][len(str2)]

    forfeit = 0 if (str1[-1] == str2[-1]) else 1
    matrix[len(str1)][len(str2)] = min(LevRecursionMatrix(str1, str2[:-1], matrix) + 1,
                                       LevRecursionMatrix(str1[:-1], str2, matrix) + 1,
                                       LevRecursionMatrix(str1[:-1], str2[:-1], matrix) + forfeit)
    return matrix[len(str1)][len(str2)]


def LevTable(str1, str2, Dam=False):
    len_i = len(str1) + 1
    len_j = len(str2) + 1
    table = [[i + j for j in range(len_j)] for i in range(len_i)]

    for i in range(1, len_i):
        for j in range(1, len_j):
            forfeit = 0 if (str1[i - 1] == str2[j - 1]) else 1
            table[i][j] = min(table[i - 1][j] + 1,
                              table[i][j - 1] + 1,
                              table[i - 1][j - 1] + forfeit)
            if Dam:
                if (i > 1 and j > 1) and str1[i - 1] == str2[j - 2] and str1[i - 2] == str2[j - 1]:
                    table[i][j] = min(table[i][j], table[i - 2][j - 2] + 1)
    if OUTPUT_MODE:
        OutputTable(table, str1, str2)
    return table[-1][-1]


def getStr():
    str1 = input("Введите первую строку: ")
    str2 = input("Введите вторую строку: ")
    return str1, str2


def getMatrixAndRun(str1="", str2=""):
    if str1 == "" and str2 == "":
        str1, str2 = getStr()
    len_i = len(str1) + 1
    len_j = len(str2) + 1
    table = [[i + j for j in range(len_j)] for i in range(len_i)]
    for i in range(len_i):
        for j in range(len_j):
            if i != 0 and j != 0:
                table[i][j] = None

    res = LevRecursionMatrix(str1, str2, table)
    if OUTPUT_MODE:
        OutputTable(table, str1, str2)
        print("Результат == ", res)


def GetStrAndRun(function, Dam=False):
    str1, str2 = getStr()
    res = function(str1, str2, Dam)
    print("Результат == ", res)


def TimeAnalysis(function, nIter, strLen=5, Dam=False):
    t1 = time.process_time()
    for i in range(nIter):
        str1 = RandomString(strLen)
        str2 = RandomString(strLen)
        function(str1, str2, Dam)
    t2 = time.process_time()
    return (t2 - t1) / nIter


def TimeAnalysisMatrix(nIter, strLen=5):
    t1 = time.process_time()
    for i in range(nIter):
        str1 = RandomString(strLen)
        str2 = RandomString(strLen)
        getMatrixAndRun(str1, str2)
    t2 = time.process_time()
    return (t2 - t1) / nIter


def PrintMenu():
    case = input("\n\n1 - Нахождение расстояния Левенштейна рекурсивно\n" +
                 "2 - Нахождение расстояния Левенштейна рекурсивно с матрицей\n" +
                 "3 - Нахождение расстояния Левенштейна матрично\n" +
                 "4 - Нахождение расстояния Дамерау - Левенштейна рекурсивно\n" +
                 "5 - Нахождение расстояния Дамерау - Левенштейна матрично\n" +
                 "6 - Сравнение алгоритмов\n" +
                 "0 - Выход\n")
    return case


def testFunctions():
    nIter = int(input("Введите кол-во итераций: "))
    strLen = int(input("Введите длину строки: "))
    print("Strlen: ", strLen)
    print(" Левенштейна рекурсивно            : ", "{0:.6f}".format(TimeAnalysis(LevRecursion, nIter, strLen)))
    print(" Левенштейна рекурсивно с матрицей : ", "{0:.6f}".format(TimeAnalysisMatrix(nIter, strLen)))
    print(" Левенштейна матрично              : ", "{0:.6f}".format(TimeAnalysis(LevTable, nIter, strLen)))
    print(" Дамерау - Левенштейна рекурсивно  : ", "{0:.6f}".format(TimeAnalysis(LevRecursion, nIter, strLen, True)))
    print(" Дамерау - Левенштейна матрично    : ", "{0:.6f}".format(TimeAnalysis(LevTable, nIter, strLen, True)))


def Menu():
    global OUTPUT_MODE
    while True:
        case = PrintMenu()
        OUTPUT_MODE = (int(case) < 6)
        if case == "1":
            GetStrAndRun(LevRecursion)
        elif case == "2":
            getMatrixAndRun()
        elif case == "3":
            GetStrAndRun(LevTable)
        elif case == "4":
            GetStrAndRun(LevRecursion, True)
        elif case == "5":
            GetStrAndRun(LevTable, True)
        elif case == "6":
            testFunctions()
        elif case == "0":
            break
        else:
            print("Неизвестная команда")


if __name__ == "__main__":
    Menu()
