#include <iostream>

using namespace std;

void swapNumbers(int &a, int &b) {
	int temp = a;
	a = b;
	b = temp;
}

int main(){

	int num1 = 5;
	int num2 = 10;

	int* ptr = &num1;
	int **ptrptr = &ptr;


	cout<<"address of pointer: "<<&ptr<<endl;
	cout<<"pointing @ : "<<ptr<<endl;
	cout<<"Content of pointer: "<<*ptr<<endl;
	cout<<"pointer to pointer: "<<ptrptr[1]<<endl;

	cout<<"Before swap"<<endl;
	cout<<"num1: "<<num1<<endl;
	cout<<"num2: "<<num2<<endl;

	swapNumbers(num1, num2);
	cout<<"After swap"<<endl;
	cout<<"num1: "<<num1<<endl;
	cout<<"num2: "<<num2<<endl;
	
	return 0;
}