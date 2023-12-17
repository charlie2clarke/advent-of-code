#include <ctype.h>

#include <cmath>
#include <fstream>
#include <functional>
#include <iostream>
#include <numeric>
#include <string>
#include <vector>

#define CONCAT_INTS(numOne, numTwo) (numOne * pow(10, static_cast<int>(log10(numTwo) + 1)) + numTwo)

static bool readFile(const std::string & fileName, std::vector<std::string> & lines)
{
  std::ifstream file{fileName};
  if (!file)
  {
    std::cerr << "Cannot open file " << fileName << std::endl;
    return false;
  }
  std::function<void()> closeStream = [&file] { file.close(); };
  std::string str;
  while (std::getline(file, str))
  {
    lines.push_back(str);
  }
  closeStream();
  return true;
}

int part1(const std::vector<std::string> & lines)
{
  std::vector<int> lineSums{};

  for (auto line : lines)
  {
    std::vector<int> digits{};
    for (auto character : line)
    {
      if (!isdigit(character))
      {
        continue;
      }
      digits.push_back(character - '0');
    }
    lineSums.push_back(CONCAT_INTS(digits.front(), digits.back()));
  }

  return std::reduce(lineSums.begin(), lineSums.end());
}

int strNum(const std::string & input)
{
  const std::vector<std::string> strNums{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"};

  for (int i = 0; i < strNums.size(); i++)
  {
    auto strNum = strNums[i];
    if (input.size() < strNum.size())
    {
      continue;
    }

    if (input.substr(0, strNum.size()) == strNum)
    {
      return i + 1;
    }
  }

  return -1;
}

int part2(const std::vector<std::string> & lines)
{
  std::vector<int> lineSums{};

  for (auto line : lines)
  {
    std::vector<int> digits{};
    for (int i = 0; i < line.size(); i++)
    {
      if (isdigit(line[i]))
      {
        digits.push_back(line[i] - '0');
        continue;
      }
      auto num = strNum(line.substr(i));
      if (num != -1)
      {
        digits.push_back(num);
      }
    }
    lineSums.push_back(CONCAT_INTS(digits.front(), digits.back()));
  }

  return std::reduce(lineSums.begin(), lineSums.end());
}

int main(int argc, char * argv[])
{
  std::vector<std::string> lines{};
  if (argc == 2)
  {
    if (!readFile(argv[1], lines))
    {
      return EXIT_FAILURE;
    }
  }

  std::cout << "Part 1: " << part1(lines) << std::endl;
  std::cout << "Part 2: " << part2(lines) << std::endl;
}