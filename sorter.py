import os 
import logging  
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

# Step 1: Import the necessary modules
# os for file operations, and logging for logging purposes
# Step 2: Get the list of files in the directory
# Step 3: Iterate over the list of files and extract the file extension using os.path.splitext
# Step 4: Create a directory for the file extension if it does not already exist
# Step 5: Move the file to the corresponding directory using os.rename
