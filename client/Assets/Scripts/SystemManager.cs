using UnityEngine;

namespace ProtoBufferExample.Client
{
    public class SystemManager : MonoBehaviour
    {
        public static SystemManager Instance { get; private set; }

        [Header("Connection Settings")]
        [SerializeField] private string _serverAddress = "127.0.0.1";
        [SerializeField] private int _serverPort = 7777;

        [SerializeField] private ConnectionType _connectionType = ConnectionType.TCP;
        [SerializeField] private SerializerType _serializerType = SerializerType.Protobuf;

        public enum ConnectionType { TCP, WebSocket }
        public enum SerializerType { Protobuf, JSON }

        // Public properties to access Model
        public Models.ConnectionModel ConnectionModel { get; private set; }
        private Presenters.ConnectionPresenter _connectionPresenter; // Reserved place for the presenter

        private void Awake()
        {
            if (Instance != null && Instance != this)
            {
                Destroy(gameObject);
                return;
            }
            Instance = this;
            DontDestroyOnLoad(gameObject);

            InitializeSystemComponents();
        }

        private void InitializeSystemComponents()
        {
            IConnection connection = null;
            ISerializer serializer = null;

            // Instantiate the chosen IConnection
            if (_connectionType == ConnectionType.TCP)
            {
                connection = new TCPClientTransport();
            }
            else if (_connectionType == ConnectionType.WebSocket)
            {
                connection = new WebSocketClientTransport();
            }

            // Instantiate the chosen ISerializer
            if (_serializerType == SerializerType.Protobuf)
            {
                serializer = new ProtobufSerializer();
            }
            else if (_serializerType == SerializerType.JSON)
            {
                serializer = new JSONSerializer();
            }

            if (connection == null || serializer == null)
            {
                Debug.LogError("Failed to initialize connection or serializer. Check SystemManager settings.");
                return;
            }

            ConnectionModel = new Models.ConnectionModel(connection, serializer);
            // _connectionPresenter is not instantiated here because it needs a concrete IConnectionView instance.
            // It will be instantiated and stored when a View calls GetConnectionPresenter.
        }

        // This method will be called by the View to register itself and get the presenter
        public Presenters.ConnectionPresenter GetConnectionPresenter(Views.IConnectionView view)
        {
            if (ConnectionModel == null)
            {
                Debug.LogError("ConnectionModel is not initialized in SystemManager.");
                return null;
            }
            // Instantiate Presenter here, as it needs the concrete view instance
            _connectionPresenter = new Presenters.ConnectionPresenter(ConnectionModel, view); // Store the created presenter
            return _connectionPresenter;
        }

        public string ServerAddress => _serverAddress;
        public int ServerPort => _serverPort;
    }
}