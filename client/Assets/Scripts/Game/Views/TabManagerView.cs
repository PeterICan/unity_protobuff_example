using UnityEngine;
using UnityEngine.UI;
using TMPro;
using System.Collections.Generic;

namespace ProtoBufferExample.Client.Game.Views
{
    public class TabManagerView : MonoBehaviour
    {
        [Header("Tab Buttons")]
        [SerializeField] private Button _connectionTabButton;
        [SerializeField] private Button _testTabButton;
        [SerializeField] private Button _gamerInfoTabButton;

        [Header("Tab Panels")]
        [SerializeField] private GameObject _connectionPanel;
        [SerializeField] private GameObject _testFunctionPanel;
        [SerializeField] private GameObject _gamerInfoPanel;

        private Button _currentActiveButton;
        private GameObject _currentActivePanel;

        private void Awake()
        {
            // Add listeners to tab buttons
            _connectionTabButton?.onClick.AddListener(() => OnTabButtonClicked(_connectionTabButton, _connectionPanel));
            _testTabButton?.onClick.AddListener(() => OnTabButtonClicked(_testTabButton, _testFunctionPanel));
            _gamerInfoTabButton?.onClick.AddListener(() => OnTabButtonClicked(_gamerInfoTabButton, _gamerInfoPanel));

            // Set initial active tab
            OnTabButtonClicked(_connectionTabButton, _connectionPanel);
        }

        private void OnTabButtonClicked(Button clickedButton, GameObject panelToActivate)
        {
            // Deactivate previous panel and button
            if (_currentActivePanel != null) _currentActivePanel.SetActive(false);
            if (_currentActiveButton != null) SetButtonVisualState(_currentActiveButton, false);

            // Activate new panel and button
            panelToActivate.SetActive(true);
            SetButtonVisualState(clickedButton, true);

            _currentActivePanel = panelToActivate;
            _currentActiveButton = clickedButton;
        }

        private void SetButtonVisualState(Button button, bool isActive)
        {
            // Example: Change button color or sprite to indicate active state
            // You might want to use a custom script for more complex visual states
            ColorBlock colors = button.colors;
            colors.normalColor = isActive ? Color.cyan : Color.white;
            button.colors = colors;

            // Optionally, change text color
            TMP_Text buttonText = button.GetComponentInChildren<TMP_Text>();
            if (buttonText != null)
            {
                buttonText.color = isActive ? Color.blue : Color.black;
            }
        }
    }
}
