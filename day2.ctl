// Day 2
// day2.txt must be in the same folder
void main()
{
  day2();
  day2(true);
}

void day2(bool part2 = false)
{
  int total = 0;
  string result;
  fileToString("day2.txt", result);
  dyn_string lines = strsplit(result, "\n");
  dyn_int ary;

  for (int i = 1; i <= dynlen(lines); i++)
  {
    string line = lines[i];

    if (strrtrim(line) == "") continue;

    dyn_string numbers = strsplit(line, " ");
    dyn_int ary;

    for (int j = 1; j <= dynlen(numbers); j++)
    {
      ary.append((int)numbers[j]);
    }

    if (isIncreasingOrDecreasing(ary))
    {
      total += 1;
    }
    else if (part2)
    {
      for (int j = 1; j <= dynlen(ary); j++)
      {
        dyn_int tmp = ary;
        dynRemove(tmp, j);

        if (isIncreasingOrDecreasing(tmp))
        {
          DebugTN("Problem Dampener success");
          total += 1;
          break;
        }
      }
    }
  }

  DebugTN(total);
}

bool isIncreasingOrDecreasing(dyn_int nums)
{
  bool isIncreasing = true;
  bool isDecreasing = true;

  for (int i = 2; i <= dynlen(nums); i++)
  {
    int diff = nums[i] - nums[i - 1];

    // Check if the difference is outside the allowed range
    if (diff < -3 || diff > 3)
    {
      return false;
    }

    // Check if the sequence is not strictly increasing
    if (diff <= 0)
    {
      isIncreasing = false;
    }

    // Check if the sequence is not strictly decreasing
    if (diff >= 0)
    {
      isDecreasing = false;
    }
  }

  return isIncreasing || isDecreasing;
}
