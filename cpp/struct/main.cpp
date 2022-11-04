#include <cstdio>
#include <fstream>
#include <string>
#include <string_view>

struct TestData final {
  int Hoge;
  char Fuga;
  float Fizz;
  short Buzz;
};

bool readTestData(std::string_view filename, TestData *testData) {
  std::string readPath(filename);
  std::ifstream file(readPath, std::ios::binary | std::ios::in);

  file.read(reinterpret_cast<char *>(testData), sizeof(*testData));

  file.close();

  return true;
}

bool writeTestData(std::string_view filename, TestData *testData) {
  std::string writePath(filename);
  std::ofstream file(writePath, std::ios::binary | std::ios::out);

  file.write(reinterpret_cast<char *>(testData), sizeof(*testData));

  file.close();

  return true;
}

int main() {
  TestData writeData{
      10,
      20,
      30,
      40,
  };

  const std::string dataName = "testdata.data";

  if (!writeTestData(dataName, &writeData)) {
    return 1;
  }

  TestData readData = {};
  if (!readTestData(dataName, &readData)) {
    return 1;
  }

  printf("writeData: {Hoge:%d Fuga:%d Fizz:%f Buzz:%d}\n", writeData.Hoge,
         writeData.Fuga, writeData.Fizz, writeData.Buzz);

  printf("readData: {Hoge:%d Fuga:%d Fizz:%f Buzz:%d}\n", readData.Hoge,
         readData.Fuga, readData.Fizz, readData.Buzz);

  return 0;
}
