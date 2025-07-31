import _ from 'lodash';

function main() {
  const numbers = [1, 2, 3, 4, 5];
  const sum = _.sum(numbers);
  console.log(`Sum of ${numbers} is ${sum}`);
  console.log("End3");
  console.log("End4");
  console.log("End5");
  console.log("End6");
  console.log("End7");
  console.log("End8");
  
  // New feature: Calculate average
  const average = _.mean(numbers);
  console.log(`Average of ${numbers} is ${average}`);
  
  // New feature: Find max and min
  const max = _.max(numbers);
  const min = _.min(numbers);
  console.log(`Max: ${max}, Min: ${min}`);
}

main();
