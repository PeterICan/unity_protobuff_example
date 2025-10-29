namespace ProtoBufferExample.Client.Game.Views
{
    public interface IConnectionView
    {
        void UpdateConnectionStatus(bool isConnected);
        void LogMessage(string message);
        // Reverted: void UpdateOtherPlayersStatus(string status);
    }
}
