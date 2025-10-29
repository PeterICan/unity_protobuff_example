using UnityEngine;
using UnityEngine.UI;
using TMPro;
using ProtoBufferExample.Client.Game.Presenters;
using ProtoBufferExample.Client.Game.Models;
using ProtoBufferExample.Client.Game.Singleton; // For SystemManager
using System.Collections.Generic; // For List

namespace ProtoBufferExample.Client.Game.Views
{
    public class ConnectionView : MonoBehaviour, IConnectionView
    {
        [Header("UI References")]
        [SerializeField] private TMP_Text _connectionStatusText;
        [SerializeField] private Button _connectButton;
        [SerializeField] private ScrollRect _logScrollRect;
        [SerializeField] private TMP_Text _otherPlayersStatusText; // Reverted to placeholder

        [Header("Log Panel Settings")]
        [SerializeField] private RectTransform _logContentParent; // The Content RectTransform of the ScrollRect
        [SerializeField] private GameObject _logMessagePrefab; // A TMP_Text prefab for individual log entries
        [SerializeField] private int _maxLogMessages = 100; // Max number of log messages to display

        private List<GameObject> _currentLogMessages = new List<GameObject>();

        private ConnectionPresenter _presenter;
        private ConnectionModel _model;

        private bool _isConnected;

        private void Start()
        {
            // Get ConnectionModel and ConnectionPresenter from SystemManager
            if (SystemManager.Instance == null)
            {
                Debug.LogError("SystemManager.Instance is null. Make sure SystemManager GameObject is in the scene and initialized.");
                return;
            }

            _model = SystemManager.Instance.ConnectionModel;
            _presenter = SystemManager.Instance.GetConnectionPresenter(this); // Pass 'this' (IConnectionView) to get the presenter

            if (_model == null || _presenter == null)
            {
                Debug.LogError("ConnectionModel or ConnectionPresenter could not be retrieved from SystemManager.");
                return;
            }

            // Wire up UI events
            _connectButton?.onClick.AddListener(
                () =>
                {
                    if (_isConnected == false)
                    {
                        _connectionStatusText.text = $"連線狀態: 連線中...";
                        _connectButton.interactable = false;
                        _presenter.Connect(SystemManager.Instance.ServerAddress, SystemManager.Instance.ServerPort);
                    }
                    else
                    {
                        _presenter.Disconnect();
                    }
                }
                );
        }

        private void OnDestroy()
        {
            // Clean up UI event listeners
            _connectButton?.onClick.RemoveAllListeners();
        }

        // IConnectionView implementation
        public void UpdateConnectionStatus(bool isConnected)
        {
            if (_connectionStatusText != null)
            {
                _connectButton.interactable = true;
                if (isConnected)
                {
                    _isConnected = true;
                    _connectButton.GetComponentInChildren<TMP_Text>().text = "斷線";
                    _connectionStatusText.text = $"連線狀態: 已連線";
                    return;
                }
                if (!isConnected)
                {
                    _isConnected = false;
                    _connectButton.GetComponentInChildren<TMP_Text>().text = "連線";
                    _connectionStatusText.text = $"連線狀態: 未連線";
                    return;
                }
            }
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

        // Reverted: UpdateOtherPlayersStatus method
        // For now, _otherPlayersStatusText will be updated directly by the presenter if needed.
        // The presenter will need to format the status string.
        public void UpdateOtherPlayersStatus(string status)
        {
            if (_otherPlayersStatusText != null)
            {
                _otherPlayersStatusText.text = status;
            }
        }
    }
}
