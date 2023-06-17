SET ROOT=%cd%
mkdir %ROOT%\..\bin\

cd %ROOT%
cd ..\cmd\rpc\user_score\&& go build  && move user_score.exe %ROOT%\..\bin\

cd %ROOT%
cd ..\cmd\rpc\chain_up\&& go build  && move chain_up.exe %ROOT%\..\bin\

cd %ROOT%
cd ..\cmd\chain_event_listener\&& go build && move chain_event_listener.exe %ROOT%\..\bin\

cd %ROOT%