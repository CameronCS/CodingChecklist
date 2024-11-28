#include <iostream>
#include <string>

std::string getstring(std::string prompt) {
    std::cout << prompt;
    std::string s;
    std::getline(std::cin, s, '\n');
    return s;
}

int main() {
    std::cout << "Welcome to the PC builder, lets build a PC\n";
    std::string cpu = getstring("Enter your CPU: ");
    std::string gpu = getstring("Enter your Grapgics Card: ");
    int ram;
    std::cout << "Enter amount of RAM: ";
    std::cin >> ram;
    std::cin.get();
    std::string motherboard = getstring("Enter Motherboard: ");
    int psu_wattage;
    std::cout << "Enter Power Supply wattage: ";
    std::cin >> psu_wattage;

    std::cout << "PC Summary\n" << "==============================================" 
              << "\nCPU: " << cpu << "\nGPU: " << gpu
              << "\nRam: " << ram << "Gb\nMotherboard: " << motherboard 
              << "\nPower Supply Wattage: " << psu_wattage << " watts"<< std::endl;
}