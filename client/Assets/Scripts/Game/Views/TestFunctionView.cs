using UnityEngine;
using ProtoBufferExample.Client.Game.Presenters;
using ProtoBufferExample.Client.Game.Models;
using ProtoBufferExample.Client.Game.Singleton;
using UnityEngine.UI; // For SystemManager

namespace ProtoBufferExample.Client.Game.Views
{
    public class TestFunctionView : MonoBehaviour, ITestFunctionView
    {
        [Header("UI References")]
        [SerializeField] private Button _C2S_PositionUpdateButton;


        private ConnectionPresenter _presenter;
        private ConnectionModel _model;

        private bool _isConnected;

        private void Start()
        {
            // // Get ConnectionModel and ConnectionPresenter from SystemManager
            // if (SystemManager.Instance == null)
            // {
            //     Debug.LogError("SystemManager.Instance is null. Make sure SystemManager GameObject is in the scene and initialized.");
            //     return;
            // }

            // _model = SystemManager.Instance.ConnectionModel;
            // _presenter = SystemManager.Instance.GetConnectionPresenter(this); // Pass 'this' (ITestFunctionView) to get the presenter

            // if (_model == null || _presenter == null)
            // {
            //     Debug.LogError("ConnectionModel or ConnectionPresenter could not be retrieved from SystemManager.");
            //     return;
            // }

            // // Wire up UI events
            // _C2S_PositionUpdateButton?.onClick.AddListener(
            //     () =>
            //     {
            //         Debug.Log("C2S_PositionUpdateButton clicked.");
            //     }
            // );
        }
    }
}