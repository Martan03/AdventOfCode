#include <iostream>
#include <vector>

int first_part();

void second_part();

char draw(int len, int x);

int main()
{
    // std::cout << "First part: " << first_part() << std::endl;
    second_part();

    return 0;
}

int first_part()
{
    std::string line;
    int loop = 0, x = 1, res = 20, sum = 0;
    
    while (getline(std::cin, line))
    {
        if (line == "noop")
        {
            ++loop;

            if (loop >= res)
            {
                std::cout << res << " " << x << std::endl;
                sum += res * x;
                res += 40;
            }

            continue;
        }

        loop += 2;
        if (loop >= res)
        {
            std::cout << res << " " << x << std::endl;
            sum += res * x;
            res += 40;
        }

        line = line.erase(0, line.find_first_of(' ') + 1);
        x += stoi(line);
    }

    return sum;
}

void second_part()
{
    std::string line;
    std::string result = "";
    std::vector<std::string> screen;
    int loop = 0, x = 1, next = 40;
    
    while (getline(std::cin, line))
    {
        ++loop;
        result += draw(result.length(), x);

        if (loop >= next)
        {
            screen.push_back(result);
            result = "";
            next += 40;
        }

        if (line == "noop")
            continue;

        ++loop;
        result += draw(result.length(), x);

        if (loop >= next)
        {
            screen.push_back(result);
            result = "";
            next += 40;
        }

        line = line.erase(0, line.find_first_of(' ') + 1);
        x += stoi(line);
    }

    for (int i = 0; i < screen.size(); ++i)
    {
        for (int j = 0; j < screen.at(i).size(); ++j)
            std::cout << screen.at(i).at(j);
        std::cout << std::endl;
    }
}

char draw(int len, int x)
{
    if (len <= x + 1 && len >= x - 1)
            return '#';
    return '.';
}
