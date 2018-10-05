let r=String.init 109 (fun i->if(i+1)mod 11==0 then '\n' else '*') in print_string r
