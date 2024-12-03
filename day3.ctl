// Day 3
// day3.txt must be in the same folder

void main()
{
  day3();
  day3part2();
}

void day3part2()
{
  bool mulEnabled = true;
  int idx = 0;
  int total = 0;
  string result;
  dyn_string numbers;
  fileToString("day3.txt", result);

  while (true)
  {
    idx = regexpSplit("(mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\))",
                      result,
                      numbers,
                      makeMapping("startPosition", idx));

    if (idx < 0) break;

    string first = numbers[1];

    if (numbers[1] == "do()")
    {
      mulEnabled = true;
    }
    else if (numbers[1] == "don't()")
    {
      mulEnabled = false;
    }
    else if (mulEnabled && first.startsWith("mul("))
    {
      dyn_string match;
      regexpSplit("mul\\((\\d+),(\\d+)\\)",
                  first,
                  match);

      if (dynlen(match) == 3)
      {
        total += ((int)match[2] * (int)match[3]);
      }
    }

    idx += numbers[1].length();
  };

  DebugTN(total);
}

void day3()
{
  int idx = 0;
  int total = 0;
  string result;
  dyn_string numbers;
  fileToString("day3.txt", result);

  while (true)
  {
    idx = regexpSplit("mul\\((\\d+),(\\d+)\\)",
                      result,
                      numbers,
                      makeMapping("startPosition", idx));

    if (idx < 0) break;

    total += ((int)numbers[2] * (int)numbers[3]);
    idx += numbers[1].length();
  };

  DebugTN(total);
}
