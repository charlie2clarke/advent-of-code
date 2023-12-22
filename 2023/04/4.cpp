#include <algorithm>
#include <array>
#include <fstream>
#include <iostream>
#include <set>
#include <sstream>
#include <string>
#include <vector>
using namespace std;

set<string> parseCard(string s)
{
  set<string> out;
  stringstream ss(s);
  string v;

  while (getline(ss, v, ' '))
  {
    if (v.empty())
      continue;

    out.emplace(v);
  }

  return out;
}

vector<array<set<string>, 2>> loadInput()
{
  ifstream file;
  file.open("input.txt");
  vector<array<set<string>, 2>> input;

  string line;
  while (getline(file, line))
  {
    line = line.substr(line.find(": ") + 2);
    size_t sep = line.find(" | ");
    set<string> wins = parseCard(line.substr(0, sep)), nums = parseCard(line.substr(sep + 3)), accWins;
    input.emplace_back(array<set<string>, 2>{{parseCard(line.substr(0, sep)), parseCard(line.substr(sep + 3))}});
  }
  file.close();

  return input;
}

size_t wins(set<string> want, set<string> got)
{
  set<string> wins;
  set_intersection(want.begin(), want.end(), got.begin(), got.end(), inserter(wins, wins.begin()));
  return wins.size();
}

int part1(vector<array<set<string>, 2>> input)
{
  int part1 = 0;

  for (const auto & card : input)
  {
    int points = 0;
    for (size_t win = 0; win < wins(card[0], card[1]); win++)
    {
      if (points == 0)
        points = 1;
      else
        points *= 2;
    }
    part1 += points;
  }

  return part1;
}

int part2(vector<array<set<string>, 2>> input)
{
  int part2 = 0, cardNum = -1;
  vector<int> copies(input.size(), 0);

  for (const auto & card : input)
  {
    cardNum++;
    int copyNum = 0;

    for (size_t win = 0; win < wins(card[0], card[1]); win++)
    {
      copyNum++;
      copies[cardNum + copyNum] += copies[cardNum] + 1;
    }
  }

  for (int & copy : copies)
    part2 += copy + 1;

  return part2;
}

int main()
{
  const auto & input = loadInput();

  cout << "Part 1: " << part1(input) << endl;
  cout << "Part 2: " << part2(input) << endl;
}