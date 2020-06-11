import java.util.Objects;

public class Solution {
    public boolean isPalindrome(int x) {
        String str = String.valueOf(x);
        int count = str.length() / 2;
        for (int i = 0; i < count; i++) {
            if (!Objects.equals(str.charAt(i), str.charAt(str.length() - 1 - i))) {
                return false;
            }
        }
        return true;
    }
}