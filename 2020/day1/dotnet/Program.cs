using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

namespace day1
{
    public class Program {

        static void Main(string[] args) {
            string[] input = File.ReadAllLines("../input");
            List<int> list = (Array.ConvertAll(input, s => Int32.Parse(s))).ToList();

            Console.WriteLine($"A: {PartA(list)}, B: {PartB(list)}");
        }

        public static int PartA(List<int> numbers) {
            foreach (int i in numbers) {
                foreach (int y in numbers) {
                    if (i + y == 2020) {
                        return i * y;
                    }
                }
            }
            return 0;
        }

        public static int PartB(List<int> numbers) {
            foreach (int i in numbers) {
                foreach (int y in numbers) {
                    foreach (int j in numbers) {
                        if (i + y + j == 2020) {
                            return i * y * j;
                        }
                    }
                }
            }
            return 0;
        }
    }
}
