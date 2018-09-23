class ReverseKeyword {
    public static void main(String[] args) {
        String testString = "TEST STRING KODLAMA PRATİKLERİ";
        System.out.println(reverse(testString, 0));
    }

    private static String reverse(String testString, int index) {
        if (index < testString.length()) {
            return reverse(testString, index + 1) + testString.charAt(index);
        } else {
            return "";
        }
    }
}