import java.util.Scanner;

public class kelimeproblem 
{

    public String donder(String kelime) 
    {     
        if ((kelime==null)||(kelime.length() <= 1) )
            return kelime;
        return donder(kelime.substring(1)) + kelime.charAt(0);
    }
    public static void main(String[] args) 
    {
        Scanner scan= new Scanner(System.in);
        Kelimeproblem obj=new Kelimeproblem();
        String kelime= scan.nextLine();
        System.out.println(kelime); 
        System.out.println(obj.reverse(kelime));
    }    
}