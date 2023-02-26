WITNESSNAME=$(grep -oP '(?<=SEE INTERVIEW #)[0-9]+' streets/Buckingham_Place)
echo $WITNESSNAME
cat interviews/interview-${WITNESSNAME}
echo $MAIN_SUSPECT