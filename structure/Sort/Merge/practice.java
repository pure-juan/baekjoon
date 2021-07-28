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
        }
        int a = (i == middle) ? j : i;
        int criteria = i == middle ? end : middle;
        for (; a < criteria; a++) {
            merged[k] = array[a];
            k++;
        }

        for (int l = start; l < end; l++) {
            array[l] = merged[l];
        }
    }

    static void split(int[] array, int start, int end) {
        if (end - start > 1) {
            split(array, start, (start + end)/2);
            split(array,  (start + end)/2, end);
        }
        merge(array, start, (start + end)/2, end);
    }

    public static void main(String[] args) {
        int[] array = new int[]{26, 32, 76, 48, 38, 96, 12, 68, 50, 94, 7, 22, 34, 10, 58, 65, 87, 97, 2, 6};
        merged = new int[array.length];

        split(array, 0, array.length);
        for (int i = 0; i < array.length; i++) {
            System.out.printf("%d ", merged[i]);
        }
    }
}