import java.io.BufferedReader;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Collections;
import java.util.Set;
import java.util.TreeSet;

/** Day 1 solution. */
public final class Main {

  /** Main. */
  public static void main(String[] args) throws IOException {
    Path path = Paths.get("input.txt");
    BufferedReader reader = Files.newBufferedReader(path);

    Set<Long> calories = new TreeSet<>(Collections.reverseOrder());
    long current = 0;

    String line;
    while ((line = reader.readLine()) != null) {
      if (line.equals("")) {
        calories.add(current);
        current = 0;
      } else {
        current += Integer.parseInt(line);
      }
    }
    calories.add(current);

    Long[] caloriesArray = calories.toArray(new Long[0]);
    System.out.println(caloriesArray[0] + caloriesArray[1] + caloriesArray[2]);
  }
}
