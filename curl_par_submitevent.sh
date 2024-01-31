seq 1 13 | xargs -n1 -P13  curl -X POST http://localhost:80/api/v1/submitevent \
                                -H 'Content-Type: application/json' \
                                -d '{"Customer": "HumanUser", "EventType": "generatingProfit", "Specifics": {"what":"light bills", "for_month": "June"}}'
