#include "stdafx.h"
#include "iostream"
#include "string"

using namespace std;

int dijitalToplam(int sayi){

	int kalan = sayi%9;
	int bolum = sayi/9;
	int kucukSayi = 0;

	if( sayi <= 9){	
		kucukSayi = 10 + (sayi-1);
	}
	else{
		string basamaklar ="";
		
		if(kalan == 0){

			for(int i =0;i<bolum; i++){
				basamaklar = basamaklar.append("1");
			}
			kucukSayi = atoi(basamaklar.c_str());
			kucukSayi = kucukSayi*9;
		}
		else{
			basamaklar = to_string(kalan);
			for(int i=0;i<bolum;i++){
				basamaklar=basamaklar.append("9");
			}

			kucukSayi =atoi(basamaklar.c_str());
		}
	}

	return kucukSayi;
}

int _tmain(int argc, _TCHAR* argv[])
{
	int deger=0;
	cout<<"Sayi Giriniz: ";
	cin>>deger;
	cout<<endl;
	cout<<dijitalToplam(deger)<<endl;		
	system("pause");
	return 0;
}
