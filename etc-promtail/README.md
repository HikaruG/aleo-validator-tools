# Mock Data for Aleo Logs Metrics  

Currently the network's not enabled to everyone.  
One can check how the Aleo Logs dashboard works by using a mock data.  


## Usage 

Put a mock-data.log containing old validator log inside the mock-data/ folder, and run:  
`chmod +x append_log.sh && ./append_log.sh`  
This will create a mock_aleo.log that the promtail will be listening.  
You can then check the Aleo Logs (Mock) dashboard to see the content.  