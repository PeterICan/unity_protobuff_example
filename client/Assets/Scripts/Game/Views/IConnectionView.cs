namespace ProtoBufferExample.Client.Game.Views
{
    public interface IConnectionView
    {
        // Connection Tab
        void SetConnectionStatusText(string text);
        void SetConnectButtonText(string text);
        void SetConnectButtonEnabled(bool enabled);
        void UpdateOtherPlayersStatus(string status);
        void LogMessage(string message);

        // Test Function Tab
        void SetTestButtonsInteractable(bool interactable);
        void ShowTestResult(string result);
    }
}
