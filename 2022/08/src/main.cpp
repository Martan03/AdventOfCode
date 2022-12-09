#include <iostream>
#include <vector>

int first_part();

int second_part();

std::vector<std::vector<int>> get_trees();

bool is_hidden(std::vector<std::vector<int>> trees, int x, int y);

int get_scenic_score(std::vector<std::vector<int>> trees, int x, int y);

int main()
{
    //std::cout << "First part: " << first_part() << std::endl;
    std::cout << "Second part: " << second_part() << std::endl;

    return 0;
}

int first_part()
{
    std::vector<std::vector<int>> trees = get_trees();

    int visible = 0;

    for (int i = 0; i < trees.size(); ++i)
        for (int j = 0; j < trees[i].size(); ++j)
            visible += !is_hidden(trees, j, i);
    
    return visible;
}

int second_part()
{
    std::vector<std::vector<int>> trees = get_trees();

    int score = 0;

    for (int i = 0; i < trees.size(); ++i)
        for (int j = 0; j < trees[i].size(); ++j)
            score = std::max(score, get_scenic_score(trees, j, i));
            
    return score;
}

std::vector<std::vector<int>> get_trees()
{
    std::vector<std::vector<int>> trees;
    std::string line;

    while(getline(std::cin, line))
    {
        std::vector<int> treeLine;
        for (auto &c : line)
            treeLine.push_back(c - '0');
        trees.push_back(treeLine);
    }

    return trees;
}

bool is_hidden(std::vector<std::vector<int>> trees, int x, int y)
{
    int xLeft = 0, xRight = trees.front().size() - 1;
    int yUp = 0, yDown = trees.front().size() - 1;

    while (trees[y][xLeft] < trees[y][x] && xLeft <= x) ++xLeft;
    while (trees[y][xRight] < trees[y][x] && xRight >= x) --xRight;
    while (trees[yUp][x] < trees[y][x] && yUp <= y) ++yUp;
    while (trees[yDown][x] < trees[y][x] && yDown >= y) --yDown;

    return xLeft < x && xRight > x && yUp < y && yDown > y;
}

int get_scenic_score(std::vector<std::vector<int>> trees, int x, int y)
{
    int xLeft = x - 1, xRight = x + 1;
    int yUp = y - 1, yDown = y + 1;
    
    while (xLeft > 0 && trees[y][xLeft] < trees[y][x]) --xLeft;
    while (xRight < trees.front().size() - 1 && trees[y][xRight] < trees[y][x]) ++xRight;
    while (yUp > 0 && trees[yUp][x] < trees[y][x]) --yUp;
    while (yDown < trees.size() - 1 && trees[yDown][x] < trees[y][x]) ++yDown;

    return (x - xLeft) * (xRight - x) * (y - yUp) * (yDown - y);
}
