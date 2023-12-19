#include <fstream>
#include <iostream>
#include <map>
#include <sstream>
#include <vector>
using namespace std;

int main()
{
  ifstream file;
  file.open("input.txt");
  map<char, int> possibleGames{{'r', 12}, {'g', 13}, {'b', 14}};
  vector<bool> games;
  string game;
  int power = 0;

  while (getline(file, game))
  {
    game = game.substr(game.find_first_of(":") + 2);
    map<char, int> maxColours{{'r', 0}, {'g', 0}, {'b', 0}};
    stringstream gameStream(game);
    string set;
    bool possible = true;

    while (getline(gameStream, set, ';'))
    {
      stringstream setStream(set);
      string cubes;

      while (getline(setStream, cubes, ','))
      {
        int num;
        char colour;

        sscanf(cubes.c_str(), "%d %c", &num, &colour);
        if (num > possibleGames[colour])
        {
          possible = false;
        }

        if (num > maxColours[colour])
        {
          maxColours[colour] = num;
        }
      }
    }
    games.push_back(possible);
    int gamePower = 1;
    for (const auto & pair : maxColours)
    {
      gamePower *= pair.second;
    }
    power += gamePower;
  }
  file.close();

  int part1 = 0;
  for (int i = 0; i < games.size(); i++)
  {
    if (games[i])
    {
      part1 += i + 1;
    }
  }

  cout << "Part 1: " << part1 << endl;
  cout << "Part 2: " << power << endl;
}
