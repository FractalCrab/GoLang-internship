#include <iostream>
#include <string.h>
using namespace std;

struct Emp
{
    string name;
    int id;
    int salary;
};
void display(Emp emp){
       cout<<"ID of employee: "<<emp.id<<endl;
       cout<<"Name of employee: "<<emp.name<<"\n";
       cout<<"Salary of employee: Rs. "<<emp.salary<<endl;

}
void higher(Emp &emp1, Emp &emp2){
   // cout<<emp1.salary;
   if(emp1.salary>emp2.salary){
       display(emp1);
     }
    else{
     display(emp2);
    }
}
int main(){
    Emp emp1,emp2;
    cin>>emp1.id;
    getline(cin,emp1.name);;
    cin>>emp1.salary>>emp2.id;
    getline(cin,emp2.name);
    cin>>emp2.salary;
    higher(emp1, emp2);
    
 
}