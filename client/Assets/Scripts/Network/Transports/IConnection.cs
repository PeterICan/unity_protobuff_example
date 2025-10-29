using System;

namespace ProtoBufferExample.Client
{
    public interface IConnection
    {
        void Connect(string address, int port);
        void Disconnect();
        void Send(byte[] data);
        event Action<byte[]> OnMessageReceived;
        event Action<bool> OnConnectionStatusChanged; // New: Notify about connection status changes
        bool IsConnected { get; } 
    }
}