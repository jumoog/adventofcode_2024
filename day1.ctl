// Day 1
// day1.txt must be in the same folder
void main()
{
  int total = 0;
  dyn_int left, right;
  string result;
  fileToString("day1.txt", result);
  dyn_string lines = strsplit(result, "\n");
  dyn_string numbers;

  for (int i = 1; i <= dynlen(lines); i++)
  {
    string line = lines[i];

    if (strrtrim(line) == "") continue;


    regexpSplit("(\\d+)\\s+(\\d+)", line, numbers);
    left.append((int)numbers[2]);
    right.append((int)numbers[3]);

  }

  left.sort();
  right.sort();

  for (int index = 1; index <= dynlen(left); index++)
  {
    total += abs(right[index] - left[index]);
  }

  DebugTN(total);
}
