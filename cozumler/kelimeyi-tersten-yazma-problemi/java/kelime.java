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
        Kelimeproblem obj=new Kelimeproblem();
        String kelime = System.out.println();
        System.out.println(kelime); 
        System.out.println(obj.reverse(kelime));
    }    
}