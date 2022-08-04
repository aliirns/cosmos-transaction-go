
#!/bin/bash

# https://github.com/aliirns

#          ,_---~~~~~----._         
#   _,,_,*^____      _____``*g*\"*, 
#  / __/ /'     ^.  /      \ ^@q   f 
# [  @f | @))    |  | @))   l  0 _/  
#  \`/   \~____ / __ \_____/    \   
#   |           _l__l_           I   
#   }          [______]           I  
#   ]            | | |            |  
#   ]             ~ ~             |  
#   |                            |   
#    |                           |   


yes | pylonsd keys export $1 --unsafe --unarmored-hex 2> file.txt && tail -n 1 file.txt
#rm file.txt