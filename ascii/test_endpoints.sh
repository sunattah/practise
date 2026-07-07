
# Exercise 2: /echo
echo -e "\n${BLUE}[Exercise 2: /echo]${NC}"
R2=$(curl -s -X POST -d "Hello Go" "$SERVER_URL/echo")
if [[ "$R2" == *"Hello Go"* ]]; then echo -e "${GREEN}✔ PASS: body echoed${NC}"; else echo -e "${RED}✘ FAIL: got '$R2'${NC}"; fi
R2G=$(curl -s -o /dev/null -w "%{http_code}" -X GET "$SERVER_URL/echo")
if [ "$R2G" == "405" ]; then echo -e "${GREEN}✔ PASS: GET blocked with 405${NC}"; else echo -e "${RED}✘ FAIL: expected 405 got $R2G${NC}"; fi
