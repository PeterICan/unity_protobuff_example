using UnityEngine;
using System.Net.Sockets;
using System.Text;
using System.Runtime.InteropServices;
using System.Threading.Tasks;
using System;

// Define the message header structure, ensuring it matches the memory layout of the Go struct.
[StructLayout(LayoutKind.Sequential, Pack = 1)]
public struct MessageHead
{
    public uint Len;   // 4 bytes, for the payload length
    public ushort Error; // 2 bytes, error code
    public byte Cmd;     // 1 byte, command
    public byte Act;     // 1 byte, action
    public ushort Index; // 2 bytes, sequence number
    public ushort Flags; // 2 bytes, flags

    public const int Size = 12;

    // Serializes the struct into a byte array in Little Endian order.
    public byte[] ToBytes()
    {
        byte[] arr = new byte[Size];
        int offset = 0;

        // Copy Len (uint, 4 bytes)
        Buffer.BlockCopy(BitConverter.GetBytes(Len), 0, arr, offset, 4);
        offset += 4;

        // Copy Error (ushort, 2 bytes)
        Buffer.BlockCopy(BitConverter.GetBytes(Error), 0, arr, offset, 2);
        offset += 2;

        // Copy Cmd (byte, 1 byte)
        arr[offset] = Cmd;
        offset += 1;

        // Copy Act (byte, 1 byte)
        arr[offset] = Act;
        offset += 1;

        // Copy Index (ushort, 2 bytes)
        Buffer.BlockCopy(BitConverter.GetBytes(Index), 0, arr, offset, 2);
        offset += 2;

        // Copy Flags (ushort, 2 bytes)
        Buffer.BlockCopy(BitConverter.GetBytes(Flags), 0, arr, offset, 2);

        // Ensure the system is Little Endian. If not, all the BitConverter arrays would need to be reversed.
        if (!BitConverter.IsLittleEndian)
        {
            // This is a simplified example; a robust implementation would reverse bytes if needed.
            Debug.LogWarning("System is Big Endian. Byte order might be incorrect.");
        }

        return arr;
    }

    // Creates a MessageHead from a byte array.
    public static MessageHead FromBytes(byte[] arr)
    {
        if (arr.Length < Size)
        {
            throw new ArgumentException("Byte array is too short to be a MessageHead");
        }

        MessageHead head = new MessageHead();
        int offset = 0;

        head.Len = BitConverter.ToUInt32(arr, offset);
        offset += 4;

        head.Error = BitConverter.ToUInt16(arr, offset);
        offset += 2;

        head.Cmd = arr[offset];
        offset += 1;

        head.Act = arr[offset];
        offset += 1;

        head.Index = BitConverter.ToUInt16(arr, offset);
        offset += 2;

        head.Flags = BitConverter.ToUInt16(arr, offset);

        return head;
    }
}

public class AntnetEchoTest : MonoBehaviour
{
    private const string Host = "127.0.0.1";
    private const int Port = 6666;

    async void Start()
    {
        Debug.Log("Starting Antnet Echo Test...");
        await ConnectAndSend();
    }

    private async Task ConnectAndSend()
    {
        try
        {
            using (TcpClient client = new TcpClient())
            {
                Debug.Log($"Connecting to {Host}:{Port}...");
                await client.ConnectAsync(Host, Port);
                Debug.Log("Connected!");

                NetworkStream stream = client.GetStream();

                // 1. Prepare the payload
                byte[] payload = Encoding.UTF8.GetBytes("Hello from Unity!");

                // 2. Prepare the header
                MessageHead header = new MessageHead
                {
                    Len = (uint)payload.Length,
                    Cmd = 1,
                    Act = 1,
                    // Other fields are 0 by default
                };

                // 3. Serialize the header
                byte[] headerBytes = header.ToBytes();

                // 4. Combine header and payload into the final packet
                byte[] packet = new byte[MessageHead.Size + payload.Length];
                Buffer.BlockCopy(headerBytes, 0, packet, 0, MessageHead.Size);
                Buffer.BlockCopy(payload, 0, packet, MessageHead.Size, payload.Length);

                // 5. Send the packet
                await stream.WriteAsync(packet, 0, packet.Length);
                Debug.Log($"Sent packet: {BitConverter.ToString(packet)}");

                // 6. Wait for the echo response
                byte[] responseBuffer = new byte[1024];
                int bytesRead = await stream.ReadAsync(responseBuffer, 0, responseBuffer.Length);
                if (bytesRead > 0)
                {
                    Debug.Log($"Received {bytesRead} bytes.");

                    // 7. Parse the response
                    if (bytesRead >= MessageHead.Size)
                    {
                        byte[] receivedHeaderBytes = new byte[MessageHead.Size];
                        Buffer.BlockCopy(responseBuffer, 0, receivedHeaderBytes, 0, MessageHead.Size);
                        MessageHead receivedHeader = MessageHead.FromBytes(receivedHeaderBytes);

                        byte[] receivedPayload = new byte[receivedHeader.Len];
                        Buffer.BlockCopy(responseBuffer, MessageHead.Size, receivedPayload, 0, (int)receivedHeader.Len);
                        string receivedMessage = Encoding.UTF8.GetString(receivedPayload);

                        Debug.Log($"<color=green>Echo success!</color>\nReceived Header: [Len={receivedHeader.Len}, Cmd={receivedHeader.Cmd}, Act={receivedHeader.Act}]\nReceived Payload: \"{receivedMessage}\"");
                    }
                }
                else
                {
                    Debug.Log("No response received.");
                }
            }
        }
        catch (Exception e)
        {
            Debug.LogError($"An error occurred: {e.Message}");
        }
    }
}
