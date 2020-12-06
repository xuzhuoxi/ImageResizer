:: IconGen.exe -size=32 -in=./icon/source/icon.png -inFolder=false -out=./icon/target/icon32x32.png -outFolder=false -format=png -ratio=75

:: IconGen.exe -size=32,64,128,256x256,512 -in=./icon/source/icon.png -inFolder=false -out=./icon/target/ -outFolder=true -format=png -ratio=75

:: IconGen.exe -size=32,64,128,256x256,512 -in=./icon/source/ -inFolder=true -out=./icon/target/ -outFolder=true -format=png -ratio=75

IconGen.exe -size=32,64,128,256x256,512 -in=./icon/source -inFolder=true -out=./icon/target/ -outFolder=true -ratio=100

pause