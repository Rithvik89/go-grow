1. Print the text of each line that appears more than once in the std input,preeceded by its count.
2. Print the count and text of lines that appear more than once in the input. It reads from the standard input or from a list of named files.
3. Use ReadFile func from ioutil package .




*io.Read is an interface defined in the io package. It is not specifically tied to reading from files but rather provides a general way to read data from any source that implements the Read method. This method reads data into a byte slice and returns the number of bytes read along with any error encountered. You can use it to read from files, network connections, memory buffers, and other sources that implement the Read method.*

*ioutil.ReadFile is a utility function provided by the ioutil package specifically for reading the entire contents of a file into memory. It abstracts away some of the lower-level details of opening and reading from a file, making it convenient for simple file reading tasks. It returns a byte slice containing the contents of the file and any error encountered.*

**In summary:**
Use io.Read when you need more control over reading data from a file or when you're working with custom data sources that implement the Read method.
Use ioutil.ReadFile for simpler cases where you just want to read the entire contents of a file into memory. It's convenient for smaller files or when you don't need fine-grained control over the reading process.