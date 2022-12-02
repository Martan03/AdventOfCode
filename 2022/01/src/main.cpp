#include <iostream>

/// @brief Checks if string is a number
/// @param text text to be checked
/// @return true if is a number, else false
bool IsStrNumber(std::string text);

/// @brief Sorts integer array
/// @param arr integer array
void SortArray(int arr[]);

// Usage: ./main <data.txt
int main()
{
    int cal = 0;
    int maxCal[] = { 0, 0, 0 };

    std::string line;
    while (getline(std::cin, line))
    {
        if (!IsStrNumber(line))
        {
            if (cal > maxCal[0])
            {
                maxCal[0] = cal;
                SortArray(maxCal);
            }
            cal = 0;
            continue;
        }

        cal += std::stoi(line);
    }

    std::cout << maxCal[0] + maxCal[1] + maxCal[2] << std::endl;
}

bool IsStrNumber(std::string text)
{
    if (text.size() <= 0)
        return false;

    for (char &c : text)
        if (!std::isdigit(c))
            return false;
    
    return true;
}

void SortArray(int arr[])
{
    for (int i = 0; i < 3; ++i)
    {
        for (int j = 1; j < 3 - i; ++j)
        {
            if (arr[j] < arr[j - 1])
            {
                int temp = arr[j];
                arr[j] = arr[j - 1];
                arr[j - 1] = temp;
            }
        }
    }
}
