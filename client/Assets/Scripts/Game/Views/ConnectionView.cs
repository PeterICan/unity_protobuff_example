using UnityEngine;
using UnityEngine.UI;
using TMPro;
using ProtoBufferExample.Client.Game.Presenters;
using ProtoBufferExample.Client.Game.Models;
using ProtoBufferExample.Client.Game.Singleton; // For SystemManager
using System.Collections.Generic;
using JsonApi; // For List

namespace ProtoBufferExample.Client.Game.Views
{
    public class ConnectionView : MonoBehaviour, IConnectionView
    {
        [Header("Connection Tab UI References")]
        [SerializeField] private TMP_Text _connectionStatusText;
        [SerializeField] private Button _connectButton;
        [SerializeField] private ScrollRect _logScrollRect;
        [SerializeField] private TMP_Text _otherPlayersStatusText; // Reverted to placeholder

        [SerializeField] private TMP_Text _connectButtonText;


        [Header("Log Panel Settings")]
        [SerializeField] private RectTransform _logContentParent; // The Content RectTransform of the ScrollRect
        [SerializeField] private GameObject _logMessagePrefab; // A TMP_Text prefab for individual log entries
        [SerializeField] private int _maxLogMessages = 100; // Max number of log messages to display

        [Header("Test Function Tab UI")]
        [SerializeField] private Button _C2S_PositionUpdateButton;
        [SerializeField] private Button _C2S_ChatMessageButton;
        [SerializeField] private Button _C2S_AttackButton;
        [SerializeField] private TMP_Text _testResultText;

        [Header("線上玩家面板")]
        [SerializeField] private GameObject _playerCardPrefab; // A TMP_Text prefab for individual player cards
        [SerializeField] private RectTransform _playerCardsContentParent; // The Content RectTransform for player cards


        private List<GameObject> _currentLogMessages = new List<GameObject>();

        private ConnectionPresenter _presenter;
        private ConnectionModel _model;

        private void Start()
        {
            if (SystemManager.Instance == null)
            {
                Debug.LogError("SystemManager.Instance is null.");
                return;
            }

            _presenter = SystemManager.Instance.GetConnectionPresenter(this);

            if (_presenter == null)
            {
                Debug.LogError("ConnectionPresenter could not be retrieved.");
                return;
            }

            SetupConnectionTabEvents();
            SetupTestFunctionTabEvents();
        }

        private void SetupConnectionTabEvents()
        {
            _connectButton?.onClick.AddListener(() => _presenter.OnConnectButtonClicked());
        }

        private void SetupTestFunctionTabEvents()
        {
            _C2S_PositionUpdateButton?.onClick.AddListener(OnPositionUpdateClicked);
            _C2S_ChatMessageButton?.onClick.AddListener(() => _presenter.SendTestMessage("ChatMessage"));
            _C2S_AttackButton?.onClick.AddListener(() => _presenter.SendTestMessage("Attack"));
        }

        private void OnDestroy()
        {
            // Clean up UI event listeners
            _connectButton?.onClick.RemoveAllListeners();
            _C2S_PositionUpdateButton?.onClick.RemoveAllListeners();
            _C2S_ChatMessageButton?.onClick.RemoveAllListeners();
            _C2S_AttackButton?.onClick.RemoveAllListeners();

            if (_presenter != null)
            {
                // _presenter.Dispose();
            }
        }

        #region IConnectionView Implementation - Connection Tab
        public void SetConnectionStatusText(string text)
        {
            if (_connectionStatusText != null)
                _connectionStatusText.text = text;
        }

        public void SetConnectButtonText(string text)
        {
            if (_connectButtonText != null)
                _connectButtonText.text = text;
        }

        public void SetConnectButtonEnabled(bool enabled)
        {
            if (_connectButton != null)
                _connectButton.interactable = enabled;
        }



        public void UpdateOtherPlayersStatus(string status)
        {
            if (_otherPlayersStatusText != null)
                _otherPlayersStatusText.text = status;
        }


        public void LogMessage(string message)
        {
            if (_logContentParent == null || _logMessagePrefab == null)
            {
                Debug.LogWarning("Log panel not configured. Logging to console: " + message);
                return;
            }

            // Remove oldest message if limit is reached
            if (_currentLogMessages.Count >= _maxLogMessages)
            {
                Destroy(_currentLogMessages[0]);
                _currentLogMessages.RemoveAt(0);
            }

            // Instantiate new log message
            GameObject logEntryGO = Instantiate(_logMessagePrefab, _logContentParent);
            TMP_Text logText = logEntryGO.GetComponent<TMP_Text>();
            if (logText != null)
            {
                logText.text = message;
            }
            _currentLogMessages.Add(logEntryGO);

            // Auto-scroll to bottom
            if (_logScrollRect != null)
            {
                LayoutRebuilder.ForceRebuildLayoutImmediate(_logContentParent);
                var textCount = _logContentParent.childCount;
                _logContentParent.sizeDelta = new Vector2(_logContentParent.sizeDelta.x, 30f * textCount);
                _logScrollRect.verticalNormalizedPosition = 0f;
            }
        }
        #endregion
        #region IConnectionView Implementation - Test Function Tab
        public void SetTestButtonsInteractable(bool interactable)
        {
            if (_C2S_PositionUpdateButton != null)
                _C2S_PositionUpdateButton.interactable = interactable;

            if (_C2S_ChatMessageButton != null)
                _C2S_ChatMessageButton.interactable = interactable;

            if (_C2S_AttackButton != null)
                _C2S_AttackButton.interactable = interactable;
        }

        public void ShowTestResult(string result)
        {
            if (_testResultText != null)
            {
                _testResultText.text = result;
            }

            // 也可以记录到日志
            LogMessage(result);
        }

        private void OnPositionUpdateClicked()
        {
            // 示例：发送随机位置
            float x = Random.Range(0f, 100f);
            float y = Random.Range(0f, 100f);
            float z = Random.Range(0f, 100f);

            _presenter.SendPositionUpdate(x, y, z);
        }

        public void UpdatePlayerCards(List<WorldPosition> players)
        {
            if (_playerCardsContentParent == null || _playerCardPrefab == null)
            {
                Debug.LogWarning("Player cards panel not configured.");
                return;
            }

            try
            {
                // Clear existing player cards
                foreach (Transform child in _playerCardsContentParent)
                {
                    Destroy(child.gameObject);
                }

                foreach (var player in players)
                {
                    var playerCard = Instantiate(_playerCardPrefab, _playerCardsContentParent);

                    // 更安全的方式：通過名稱或標籤查找子物件
                    var playerIdText = playerCard.transform.Find("playerIdText")?.GetComponent<TMP_Text>();
                    var positionText = playerCard.transform.Find("playerPosText")?.GetComponent<TMP_Text>();

                    if (playerIdText != null)
                        playerIdText.text = $"玩家ID: {player.PlayerId}";

                    if (positionText != null)
                        positionText.text = $"位置: ({player.X:F2}, {player.Y:F2}, {player.Z:F2})";
                }
            }
            catch (System.Exception ex)
            {
                Debug.LogError($"Error updating player cards: {ex.Message}");
            }
        }

        #endregion
    }
}
