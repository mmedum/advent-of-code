import java.io.BufferedReader;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

/** Day2. */
public class Main {

  /**
   * A = Rock, B = Paper, C = Scissors, X = Rock, Y = Paper, Z = Scissors, Lose = 0, Draw = 3, Win =
   * 6.
   */
  public static void main(String[] args) throws IOException {
    Path path = Paths.get("input.txt");
    BufferedReader reader = Files.newBufferedReader(path);

    int points = 0;
    String line;
    while ((line = reader.readLine()) != null) {
      String[] choices = line.split(" ");
      // Part 1
      // points += getPointsForBattle(choices[0], choices[1]) +
      // getPointsForChoice(choices[1]);

      // Part 2
      points += getBattlePoints(choices[0], choices[1]);
    }

    System.out.println(points);
  }

  private static int getBattlePoints(String opChoice, String condition) {
    String choice = "";
    int points = 0;
    if (condition.equalsIgnoreCase("X")) {
      points += 0;
      if (opChoice.equalsIgnoreCase("A")) {
        choice = "Z";
      } else if (opChoice.equalsIgnoreCase("B")) {
        choice = "X";
      } else {
        choice = "Y";
      }
    } else if (condition.equalsIgnoreCase("Y")) {
      points += 3;
      if (opChoice.equalsIgnoreCase("A")) {
        choice = "X";
      } else if (opChoice.equalsIgnoreCase("B")) {
        choice = "Y";
      } else {
        choice = "Z";
      }
    } else {
      points += 6;
      if (opChoice.equalsIgnoreCase("A")) {
        choice = "Y";
      } else if (opChoice.equalsIgnoreCase("B")) {
        choice = "Z";
      } else {
        choice = "X";
      }
    }

    return points + getPointsForChoice(choice);
  }

  // Part1
  private static int getPointsForBattle(String op, String me) {
    if ((op.equalsIgnoreCase("A") && me.equalsIgnoreCase("X"))
        || (op.equalsIgnoreCase("B") && me.equalsIgnoreCase("Y"))
        || (op.equalsIgnoreCase("C") && me.equalsIgnoreCase("Z"))) {
      return 3;
    } else if ((op.equalsIgnoreCase("A") && me.equalsIgnoreCase("Y"))
        || (op.equalsIgnoreCase("B") && me.equalsIgnoreCase("Z"))
        || (op.equalsIgnoreCase("C") && me.equalsIgnoreCase("X"))) {
      return 6;
    } else {
      return 0;
    }
  }

  private static int getPointsForChoice(String choice) {
    if (choice.equals("X")) {
      return 1;
    } else if (choice.equals("Y")) {
      return 2;
    } else {
      return 3;
    }
  }
}
