# av-log-parser
A small project for parsing AV raw logs; one implementation in python, the other in go

Rationale: AV can be very wish-washy when it comes to exporting data. In the brief time 
since i've started heavily using it at work, I've encountered a number of issues across deployments.
Sometimes csv reports won't export if you have certain column fields added. Othertimes, your report
appears to be generated but never is. Whatever the case, it was good to have a backup option since
most of the time, downloading raw logs is an option that works. 

The algorithm I used should be good for just about any data field given in the raw log. The odd cases
are the integer cases where there are no quotations surrounding the value, in which case some
adjustments are needed to get the right slice of the log. 

More features will be added in time. 
