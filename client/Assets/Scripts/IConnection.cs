using System;

namespace ProtoBufferExample.Client
{
    public interface IConnection
    {
        void Connect(string address, int port);
        void Disconnect();
        void Send(byte[] data);
        event Action<byte[]> OnMessageReceived;
        bool IsConnected { get; }
    }
}
