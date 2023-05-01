# jobSeeker
Tool to track job contacts and status

## Basic usage

```bash
./jobSeeker -h
./jobSeeker tool.
Job add format
	"job;title;active(y/n);notes;company;url;pocId;firstContact;lastContact"
POC add format
	"poc;name;notes;company;email;phone;firstContact;lastContact"

Usage:
  -add string
    	Double quoted line to add
  -jobs
    	Look at jobs
  -pocs
    	Look at POCs
  -search string
    	String to search for
```

## Use Cases

I keep seeing job postings that look great until I dig into the description. Then I see they have a focus on X, and Y technologies, and I don't know those at all. I needed a way to track my notes on jobs, to see if I had already applied there, or if I needed to follow up, or if the skills/salary/location wasn't a good fit. This tool lets you search jobs for key words, like the Company name or a specific job requisition number:

```shell
./jobSeeker --search MyCo --jobs
```

From my sample dataset, this gives a result of:

```shell
4;Cyber Python SW Eng;n;Python, Linux, AWS;MyCo;;1;20230501;20230501;
```

If the job I'm looking at seems similar, then I know if I've applied for that job. The notes field, "Python, Linux, AWS", tells me about the job. I might add a note about "salary too low", or "not remote". The field before that, "n", tells me if this job is still on my active list. Either it's not a good fit, or I applied but got the "thanks, but no thanks", email.


Some years ago I wrote a shell program to go through my contacts list and send an email saying I was on the job market. I actually got a job from that email, and this tool also tracks Points of Contact. It doesn't yet write that email, but you can get a start with:

```shell
./jobSeeker --pocs |awk -F";" '{print $5,"  ",  $2}'
```

Which gives something like:

```shell
jsmythe@myco.com    John Smythe
jzan@acme.com    Jane Zan
f.stone@example.com    Frederick F. Stone
```


### Add a Point of Contact
A Point of Contact (POC) is a specific person you connect with for your job search. They might be a recruiter for a hiring agency, or someone in the company you're interested in. This tool tracks basic info, you should connect via LinkedIn or similar so that you can get a well rounded understanding of the person. 

First, add the POC. Note that the two dates are left off the input string, they are auto-generated. Other fields that are empty or missing get an empty string value.

```bash
./jobSeeker --add "poc;John Smythe;quick responder;MyCo;jsmythe@myco.com; 555.1212"
```

This adds a line like this to the jobSeeker_pocs.data file. The first element, "1" in this case, is the POC id number you want.

```
1;John Smythe;quick responder;MyCo;jsmythe@myco.com;555.1212;20230501;20230501;;
```

### Add a Job
A job can use the POC identifier, if you have one. Seach for the POC:

```bash
./jobSeeker --search "John Smythe" --pocs
```

This gives you the line(s) you searched for. Choose the POC id number you want. Now add the job data, using double quotes (") to prevent the shell from interpreting the semi-colons (;) as command ends.

```bash
./jobSeeker -add "job; Cyber Python SW Eng; n; Cyber Python SW Eng;  MyCo ;;1"
```

This gives you a line similar to:

```bash
4;Cyber Python SW Eng;n;Cyber Python SW Eng;MyCo;;1;20230501;20230501;
```



## TODO
Lots of improvements are planned, beyond just making the code better. We will have a way to deactivate a job, showing that it is no longer an option for us, as well as searching all jobs, or just active/inactive ones. Other plans include printing the Job data linked to the POC data, if provided, and searching jobs/pocs that haven't been updated in a while. That sort of follow-up often leads to good things!

