

//判断闰年
#include <stdio.h>
int main()
{

    //声明一个变量来保存用户输入的信息
    int year

    printf("请输人年份：\n");
    //从控制台接受用户输入的信息
    scanf("%d",&year);

    //因为用户输入的信息保存到year变量,从year变量中来判断
    //这一步来进行判断    该年份能被400整除。
    if(year%400==0)//此处为判断条件,只有为真才进入
       //进到这个里面
       printf("%d 此年是闰年\n",year);
    else //此处为假进入
    {
        //接着判断是不是闰年的第二个条件该年份能被 4 整除同时不能被 100 整除,如果等于0
        if(year%4 == 0 && year%100 != 0)
          //是闰年进入这个代码块
           printf("%d 此年是闰年\n",year);
        else
         //不就是闰年进入这里
          printf("%d 此年非闰年\n",year);
    }
    return 0;
}




#include <stdio.h>
int main()
{
   //声明变量
    int hun, ten, ind, i;
    printf("result is:");
    for( i=100; i<1000; i++ )  /*整数的取值范围*/
    {
    //利用for循环控制100-999个数，每个数分解出个位，十位，百位。
        hun = i / 100;
        ten = (i-hun*100) / 10;
        ind = i % 10;
        if(i == hun*hun*hun + ten*ten*ten + ind*ind*ind)  /*各位上的立方和是否与原数n相等*/
            printf("%d  ", i);
    }
    printf("\n");

    return 0;
}