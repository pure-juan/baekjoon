package kr.update;

public class Main {

    static int[] merged;

    // [7, 6] [5, 8]
    // 합치는 넘
    static void merge(int[] array, int start, int middle, int end) {

        int i = start;
        int j = middle;

        int k = start;
        while (i < middle && j < end) {
            if (array[i] < array[j]) {
                merged[k] = array[i];
                i++;
            } else {
                merged[k] = array[j];
                j++;
            }
            k++;
//            for (int a = 0; a < array.length; a++) {
//                System.out.printf("%d ", merged[a]);
//            }
//            System.out.println();
        }
//        System.out.println(i + " " + j);
        // 남아있는 놈 다 넣기
        for (int a = (i == middle) ? j : i; a < (i == middle ? end : middle); a++) {
            System.out.println(a);
            merged[k] = array[a];
            k++;
        }

        for (int l = start; l < end; l++) {
            array[l] = merged[l];
        }
    }
    
    static void split(int[] array, int start, int end) {
        while (start < end) {
            split(array, start, (start + end)/2);
            split(array,  (start + end)/2 + 1, end);
            merge(array, start, (start + end)/2, end);
        }
    }

    public static void main(String[] args) {
        int[] array = new int[]{7, 6, 5, 8};
        merged = new int[array.length];

        split(array, 0, array.length);
        for (int i = 0; i < array.length; i++) {
            System.out.printf("%d ", merged[i]);
        }
    }
}
