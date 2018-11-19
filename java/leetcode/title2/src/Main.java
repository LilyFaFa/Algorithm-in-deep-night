
import java.util.HashMap;

public class Main {

    public static void main(String[] args) {
        System.out.println("Hello World!");
    }


    public static int test(int n, int k, int t, int a[]) {
        int i = 0;
        int r = a.length - 1;
//        int l = 0;
        for (; r > 0; r--) {
            for (int l = 0; l <= r; l++) {
                HashMap<Integer, Integer> map = new HashMap<Integer, Integer>();
                if (r - l + 1 == k) {
                    for (int m = l; m <= r; m++) {
                        if (map.get(a[m]) == null) {
                            map.put(a[m], 1);
                        } else {
                            map.put(a[m], map.get(a[m]) + 1);
                        }
                    }
                    for (int mapkey : map.keySet()) {
                        if (map.get(mapkey).intValue() >= t) {
                            i++;
                        }
                    }
                }
            }
        }

        return i;
    }
}