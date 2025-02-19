package org.renuka.example;

import java.util.HashMap;

public class Main {
    static HashMap<String, String> map = new HashMap<>();
    public static void main(String[] args) throws InterruptedException {
        System.out.println("Hello World");
        for (int i = 0; i < 100000000; i++) {
            if (i % 1000 == 0) {
                System.out.println("i = " + i);
                Thread.sleep(100);
            }
            map.put("key" + i, "value" + i);
        }
        Thread.sleep(100000000);
    }
}
