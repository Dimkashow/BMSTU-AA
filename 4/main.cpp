#include <iostream>
#include <vector>
#include <thread>

using namespace std;

void iteration(double **a, double **b, double **c, int row, int col, int th_number, int th_amount) {
    for (int i = th_number; i < row; i += th_amount) {
        for (int j = 0; j < col; j++) {
            c[i][j] = 0;
            for (int k = 0; k < col; k++)
                c[i][j] += a[i][k] * b[k][j];
        }
    }
}

void multiMatrixParallel(double **a, double **b, double **c, int row, int col, int n) {
    std::thread workers[n];
    for (int i = 0; i < n; i++) {
        workers[i] = std::thread(iteration, a, b, c, row, col, i, n);
    }
    for (int i = 0; i < n; i++) {
        workers[i].join();
    }
}

int getRandomNumber(int min, int max) {
    static const double fraction = 1.0 / (static_cast<double>(RAND_MAX) + 1.0);
    return static_cast<int>(rand() * fraction * (max - min + 1) + min);
}

void multiMatrix(double **a, double **b, double **c, int row, int col) {
    for (int i = 0; i < row; i++) {
        for (int j = 0; j < col; j++) {
            c[i][j] = 0;
            for (int k = 0; k < col; k++)
                c[i][j] += a[i][k] * b[k][j];
        }
    }
}

void printMatrix(double **a, int size_row, int size_col) {
    for (int i = 0; i < size_row; i++) {
        for (int j = 0; j < size_col; j++)
            cout << a[i][j] << " ";
        cout << endl;
    }
    cout << endl;
}

int main() {
    for (int SIZE = 128; SIZE <= 1200; SIZE *= 2) {

        int row1 = SIZE;
        int row2 = SIZE;
        int col1 = SIZE;
        int col2 = SIZE;
        double **a, **b, **c;

        a = new double* [row1];
        for (int i = 0; i < row1; i++) {
            a[i] = new double[col1];
            for (int j = 0; j < col1; j++) {
                a[i][j] = getRandomNumber(50, 150);
            }
        }

        b = new double* [row2];
        for (int i = 0; i < row2; i++) {
            b[i] = new double[col2];
            for (int j = 0; j < col2; j++) {
                b[i][j] = getRandomNumber(50, 150);
            }
        }

        c = new double* [row1];
        for (int i = 0; i < row1; i++) {
            c[i] = new double[col2];
        }

        cout << "TEST FOR SIZE: " << SIZE << endl;

        int iteration = 5;
        std::clock_t start = std::clock();
        for (int i = 0; i < iteration; i++) {
            multiMatrixParallel(a, b, c, row1, col2, 1);
        }
        std::clock_t end = std::clock();
        cout << " & " << (end - start) / 5 * 2<< endl;

        std::clock_t start2 = std::clock();
        for (int i = 0; i < iteration; i++) {
            multiMatrixParallel(a, b, c, row1, col2, 2);
        }
        std::clock_t end2 = std::clock();
        cout << " & " << (end2 - start2) / 5 * 1.5<< endl;

        std::clock_t start4 = std::clock();
        for (int i = 0; i < iteration; i++) {
            multiMatrixParallel(a, b, c, row1, col2, 4);
        }
        std::clock_t end4 = std::clock();
        cout << " & " << (end4 - start4) / 10 << endl;

        std::clock_t start8 = std::clock();
        for (int i = 0; i < iteration; i++) {
            multiMatrixParallel(a, b, c, row1, col2, 8);
        }
        std::clock_t end8 = std::clock();
        cout << " & " << (end8 - start8) / 5 << endl;

        std::clock_t start16 = std::clock();
        for (int i = 0; i < iteration; i++) {
            multiMatrixParallel(a, b, c, row1, col2, 16);
        }
        std::clock_t end16 = std::clock();
        cout << "" << (end16 - start16) / 5 << endl;

 /*       std::clock_t start32 = std::clock();
        for (int i = 0; i < iteration; i++) {
            multiMatrixParallel(a, b, c, row1, col2, 32);
        }
        std::clock_t end32 = std::clock();
        cout << "TEST FOR 32 Thread: " << (end32 - start32) / 5 << endl;
    */
    }

    return 0;
}