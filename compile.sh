echo "Make sure to have changed the ip and port in the code"
echo -n "Compile for linux or windows [l,w]: "
read os

if [ $os = "l" ]
then
    GOOS=linux GOARCH=amd64 go build -o hello ./rev_linux.go

elif [ $os = "w" ]
then
    echo -n "Architecture [64,86]: "
    read arch
    if [ $arch = "64" ]
    then
        GOOS=windows GOARCH=amd64 go build -o hello.exe ./rev_windows.go
    elif [ $arch = "86" ]
    then
        GOOS=windows GOARCH=386 go build -o hello.exe ./rev_windows.go
    fi
else
    echo "Invalid input"
    exit
fi

echo "done!"
