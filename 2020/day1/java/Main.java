import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class Main {

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new FileReader("../input"));
        List<Integer> list = new ArrayList<>();
        String line = br.readLine();
        while (line != null) {
            list.add(Integer.parseInt(line));
            line = br.readLine();
        }
        br.close();

        System.out.println("A: " + partA(list) + ", B: " + partB(list));
    }

    public static int partA(List<Integer> list) {
        for (int i = 0; i < list.size(); i++) {
            for (int j = i + 1; j < list.size(); j++) {
                int first = list.get(i);
                int second = list.get(j);
                if (first + second == 2020) {
                    return first * second;
                }
            }
        }
        return 0;
    }

    public static int partB(List<Integer> list) {
        for (int i = 0; i < list.size(); i++) {
            for (int j = i + 1; j < list.size(); j++) {
                for (int k = j + 1; k < list.size(); k++) {
                    int first = list.get(i);
                    int second = list.get(j);
                    int third = list.get(k);
                    if (first + second + third == 2020) {
                        return first * second * third;
                    }
                }
            }
        }
        return 0;
    }
}
