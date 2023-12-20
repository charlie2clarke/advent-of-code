#include <cmath>
#include <fstream>
#include <iostream>
#include <map>
#include <vector>
using namespace std;

struct point_t
{
  int x, y;

  bool operator<(const point_t & other) const { return std::tie(x, y) < std::tie(other.x, other.y); }

  bool operator==(const point_t & other) const { return x == other.x && y == other.y; }
};

vector<vector<char>> loadInput()
{
  ifstream file;
  file.open("input.txt");

  vector<vector<char>> input;
  string line;
  while (getline(file, line))
  {
    vector<char> chars;
    for (const char & c : line)
      chars.push_back(c);

    input.push_back(chars);
  }

  file.close();
  return input;
}

bool isValidMove(vector<vector<char>> input, int y, int x, point_t dir)
{
  int newY = y + dir.y, newX = x + dir.x;
  if (newY < 0 || newY >= input.size())
    return false;
  if (newX < 0 || newX >= input[y].size())
    return false;
  return true;
}

int concatenateIntegers(const vector<int> & numbers)
{
  int result = 0;
  for (int digit : numbers)
    result = result * 10 + digit;

  return result;
}

int main()
{
  auto input = loadInput();
  int part1 = 0, part2 = 0;
  map<point_t, vector<int>> stars;

  for (int row = 0; row < input.size(); row++)
    for (int col = 0; col < input[row].size(); col++)
    {
      auto c = input[row][col];
      if (!isdigit(c))
        continue;

      vector<int> digits;
      int moves = 0;
      bool partNum = false;
      point_t starPos{-1, -1};
      for (int lookAhd = col; lookAhd < input[row].size(); lookAhd++)
      {
        c = input[row][lookAhd];
        if (!isdigit(c))
        {
          col += moves;
          break;
        }
        moves++;

        for (const auto & dir : vector<point_t>{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}})
        {
          if (!isValidMove(input, row, lookAhd, dir))
            continue;

          point_t checkPos{lookAhd + dir.x, row + dir.y};
          auto check = input[checkPos.y][checkPos.x];
          if (!isdigit(check) && check != '.')
            partNum = true;
          if (check == '*')
            starPos = checkPos;
        }

        digits.emplace_back(c - '0');
      }

      auto num = concatenateIntegers(digits);
      if (partNum)
        part1 += num;
      if (starPos.x != -1 && starPos.y != -1)
        stars[starPos].push_back(num);
    }

  for (const auto & kv : stars)
    if (kv.second.size() == 2)
      part2 += kv.second[0] * kv.second[1];

  cout << "Part 1: " << part1 << endl;
  cout << "Part 2: " << part2 << endl;
}
