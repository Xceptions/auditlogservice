# echo '{"customer": "John", "event": "payingDues", "specifics": {"what":"light bills", "for_month": "June"}}' | nc localhost 8952  &
# echo '{"customer": "Jennifer", "event": "boughtCodingCourse", "specifics": {"platform":"Coursera", "price": "50$"}}' | nc localhost 8952  &
# echo '{"customer": "Chima", "event": "subscribedToTV", "specifics": {"channel":"Disney", "for_month": "December"}}' | nc localhost 8952  &
# echo '{"customer": "John", "event": "payingDues", "specifics": {"what":"light bills", "for_month": "June"}}' | nc localhost 8952  &
# echo '{"customer": "Jennifer", "event": "boughtCodingCourse", "specifics": {"platform":"Coursera", "price": "50$"}}' | nc localhost 8952

# curl -X POST -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJ1c2VyIjoidXNlZnVsVXNlcjMifQ.7VsYcrhAHt0rsEoaW1zi9gei4IzaZcHKiCeow_A7ckg" -d '{"customer": "Steve", "eventtype": "payingDues", "specifics": {"what":"light bills", "for_month": "June"}}' http://localhost:4000/api/v1/submitevent
# curl -X POST -d '{"customer": "Jennifer", "eventtype": "boughtCodingCourse", "specifics": {"platform":"Coursera", "price": "50$"}}' http://localhost:8952
# curl -X POST -d '{"customer": "Chima", "eventtype": "subscribedToTV", "specifics": {"channel":"Disney", "for_month": "December"}}' http://localhost:8952
# curl -X POST -d '{"customer": "Steve", "eventtype": "payingDues", "specifics": {"what":"light bills", "for_month": "June"}}' http://localhost:8952
# curl -X POST -d '{"customer": "Jennifer", "eventtype": "boughtCodingCourse", "specifics": {"platform":"Coursera", "price": "50$"}}' http://localhost:8952
# curl -X POST -d '{"customer": "Chima", "eventtype": "subscribedToTV", "specifics": {"channel":"Disney", "for_month": "December"}}' http://localhost:8952
# curl -X POST -d '{"customer": "Steve", "eventtype": "payingDues", "specifics": {"what":"light bills", "for_month": "June"}}' http://localhost:8952
# curl -X POST -d '{"customer": "Jennifer", "eventtype": "boughtCodingCourse", "specifics": {"platform":"Coursera", "price": "50$"}}' http://localhost:8952
# curl -X POST -d '{"customer": "Chima", "eventtype": "subscribedToTV", "specifics": {"channel":"Disney", "for_month": "December"}}' http://localhost:8952
# curl -X POST -d '{"customer": "Steve", "eventtype": "payingDues", "specifics": {"what":"light bills", "for_month": "June"}}' http://localhost:8952
# curl -X POST -d '{"customer": "Jennifer", "eventtype": "boughtCodingCourse", "specifics": {"platform":"Coursera", "price": "50$"}}' http://localhost:8952
# curl -X POST -d '{"customer": "Chima", "eventtype": "subscribedToTV", "specifics": {"channel":"Disney", "for_month": "December"}}' http://localhost:8952
# curl -X POST -d '{"customer": "Steve", "eventtype": "payingDues", "specifics": {"what":"light bills", "for_month": "June"}}' http://localhost:8952
# curl -X POST -d '{"customer": "Jennifer", "eventtype": "boughtCodingCourse", "specifics": {"platform":"Coursera", "price": "50$"}}' http://localhost:8952
# curl -X POST -d '{"customer": "Chima", "eventtype": "subscribedToTV", "specifics": {"channel":"Disney", "for_month": "December"}}' http://localhost:8952

# seq 1 13 | xargs -n1 -P13  curl -X POST -d '{"customer": "Steve", "eventtype": "payingDues", "specifics": {"what":"light bills", "forMonth": "June"}}' http://localhost:8952

seq 1 13 | xargs -n1 -P13  curl -X POST http://localhost:80/api/v1/submitevent \
                                -H 'Content-Type: application/json' \
                                -d '{"customer": "OdogwuPlentyFromContainer", "eventtype": "payingDues", "specifics": {"what":"light bills", "for_month": "June"}}'







#   --next \
# -d '{"customer": "Jennifer", "event": "boughtCodingCourse", "specifics": {"platform":"Coursera", "price": "50$"}}' http://127.0.0.1:8952
