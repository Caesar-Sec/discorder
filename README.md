# discorder
Go script for sending messages to discord via webhooks, with standard markdown for discord.


# Install

- ``` git clone https://github.com/Caesar-Sec/discorder.git ```
- ``` cd discorder ```
-  Edit discorder.go with your prefered editor, from here you have two options:  
    - Paste your webhook url on line 34.  
  OR  
    - Comment out line 34 and remove the // from line 32, then save and exit the editor.
    - Next run the following command and add your webhook: ```export D_NOTIFICATION_WH=[your webhook url]```
    
- ```go install discorder.go```
    

# Basic Usage
### Standard message
* ```discorder test```  

### Bold message
* ```discorder -b test```  

### Inline-code message
* ```discorder -i test```  

### Code block message
* ```discorder -c test```  

### Including file content
* ```discorder $(< test.txt)```  

# Notes
- Only one flag should be called at a time, more than this breaks the output by sending the second flag through the webhook as part of the message.
- Currently this only tool only allows for one webhook url to be used, there are ways round this issue but I plan to update it in the future to allow for more webhooks to be added. 

# To-Do
- [ ] Read webhooks from a configuration file and allow the user to select which webhook the tool should use.
- [ ] Implement a better method of including file content without first assgining it to a variable.
- [ ] Allow the output of commands to be piped to discorder.


# Credit
The original code was taken from this blog post on [Recon with discord webhooks](https://ryanwise.me/blog/recon-upgrade-discord/). I then modified it to allow for formatting.
