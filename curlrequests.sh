# echo -n "{'customer': 'John', 'event': 'payingDues', 'specifics': {'what':'light bills', 'for_month': 'June'}}" | nc localhost 8952  &
# echo -n "{'customer': 'Jennifer', 'event': 'boughtCodingCourse', 'specifics': {'platform':'Coursera', 'price': '50$'}}" | nc localhost 8952  &
# echo -n "{'customer': 'Chima', 'event': 'subscribedToTV', 'specifics': {'channel':'Disney', 'for_month': 'December'}}" | nc localhost 8952  

curl -X POST -d "{'customer': 'Steve', 'event': 'payingDues', 'specifics': {'what':'light bills', 'for_month': 'June'}}" http://localhost:8952
    --next \
    -d "'customer': 'Jennifer', 'event': 'boughtCodingCourse', 'specifics': {'platform':'Coursera', 'price': '50$'}}" http://127.0.0.1:8952
