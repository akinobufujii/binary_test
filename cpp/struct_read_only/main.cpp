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

int main() {
  const std::string dataName = "testdata.data";

  TestData readData = {};
  if (!readTestData(dataName, &readData)) {
    return 1;
  }

  printf("readData: {Hoge:%d Fuga:%d Fizz:%f Buzz:%d}\n", readData.Hoge,
         readData.Fuga, readData.Fizz, readData.Buzz);

  return 0;
}
