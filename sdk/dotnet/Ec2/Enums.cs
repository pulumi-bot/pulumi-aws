// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.Aws.Ec2
{
    [EnumType]
    public readonly struct InstancePlatform : IEquatable<InstancePlatform>
    {
        private readonly string _value;

        private InstancePlatform(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InstancePlatform LinuxUnix { get; } = new InstancePlatform("Linux/UNIX");
        public static InstancePlatform RedHatEnterpriseLinux { get; } = new InstancePlatform("Red Hat Enterprise Linux");
        public static InstancePlatform SuseLinux { get; } = new InstancePlatform("SUSE Linux");
        public static InstancePlatform Windows { get; } = new InstancePlatform("Windows");
        public static InstancePlatform WindowsWithSqlServer { get; } = new InstancePlatform("Windows with SQL Server");
        public static InstancePlatform WindowsWithSqlServerEnterprise { get; } = new InstancePlatform("Windows with SQL Server Enterprise");
        public static InstancePlatform WindowsWithSqlServerStandard { get; } = new InstancePlatform("Windows with SQL Server Standard");
        public static InstancePlatform WindowsWithSqlServerWeb { get; } = new InstancePlatform("Windows with SQL Server Web");

        public static bool operator ==(InstancePlatform left, InstancePlatform right) => left.Equals(right);
        public static bool operator !=(InstancePlatform left, InstancePlatform right) => !left.Equals(right);

        public static explicit operator string(InstancePlatform value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InstancePlatform other && Equals(other);
        public bool Equals(InstancePlatform other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct InstanceType : IEquatable<InstanceType>
    {
        private readonly string _value;

        private InstanceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InstanceType A1_2XLarge { get; } = new InstanceType("a1.2xlarge");
        public static InstanceType A1_4XLarge { get; } = new InstanceType("a1.4xlarge");
        public static InstanceType A1_Large { get; } = new InstanceType("a1.large");
        public static InstanceType A1_Medium { get; } = new InstanceType("a1.medium");
        public static InstanceType A1_XLarge { get; } = new InstanceType("a1.xlarge");
        public static InstanceType C3_2XLarge { get; } = new InstanceType("c3.2xlarge");
        public static InstanceType C3_4XLarge { get; } = new InstanceType("c3.4xlarge");
        public static InstanceType C3_8XLarge { get; } = new InstanceType("c3.8xlarge");
        public static InstanceType C3_Large { get; } = new InstanceType("c3.large");
        public static InstanceType C3_XLarge { get; } = new InstanceType("c3.xlarge");
        public static InstanceType C4_2XLarge { get; } = new InstanceType("c4.2xlarge");
        public static InstanceType C4_4XLarge { get; } = new InstanceType("c4.4xlarge");
        public static InstanceType C4_8XLarge { get; } = new InstanceType("c4.8xlarge");
        public static InstanceType C4_Large { get; } = new InstanceType("c4.large");
        public static InstanceType C4_XLarge { get; } = new InstanceType("c4.xlarge");
        public static InstanceType C5_18XLarge { get; } = new InstanceType("c5.18xlarge");
        public static InstanceType C5_2XLarge { get; } = new InstanceType("c5.2xlarge");
        public static InstanceType C5_4XLarge { get; } = new InstanceType("c5.4xlarge");
        public static InstanceType C5_9XLarge { get; } = new InstanceType("c5.9xlarge");
        public static InstanceType C5_Large { get; } = new InstanceType("c5.large");
        public static InstanceType C5_XLarge { get; } = new InstanceType("c5.xlarge");
        public static InstanceType C5a_12XLarge { get; } = new InstanceType("c5a.12xlarge");
        public static InstanceType C5a_16XLarge { get; } = new InstanceType("c5a.16xlarge");
        public static InstanceType C5a_2XLarge { get; } = new InstanceType("c5a.2xlarge");
        public static InstanceType C5a_24XLarge { get; } = new InstanceType("c5a.24xlarge");
        public static InstanceType C5a_4XLarge { get; } = new InstanceType("c5a.4xlarge");
        public static InstanceType C5a_8XLarge { get; } = new InstanceType("c5a.8xlarge");
        public static InstanceType C5a_Large { get; } = new InstanceType("c5a.large");
        public static InstanceType C5a_XLarge { get; } = new InstanceType("c5a.xlarge");
        public static InstanceType C5d_18XLarge { get; } = new InstanceType("c5d.18xlarge");
        public static InstanceType C5d_2XLarge { get; } = new InstanceType("c5d.2xlarge");
        public static InstanceType C5d_4XLarge { get; } = new InstanceType("c5d.4xlarge");
        public static InstanceType C5d_9XLarge { get; } = new InstanceType("c5d.9xlarge");
        public static InstanceType C5d_Large { get; } = new InstanceType("c5d.large");
        public static InstanceType C5d_XLarge { get; } = new InstanceType("c5d.xlarge");
        public static InstanceType C5n_18XLarge { get; } = new InstanceType("c5n.18xlarge");
        public static InstanceType C5n_2XLarge { get; } = new InstanceType("c5n.2xlarge");
        public static InstanceType C5n_4XLarge { get; } = new InstanceType("c5n.4xlarge");
        public static InstanceType C5n_9XLarge { get; } = new InstanceType("c5n.9xlarge");
        public static InstanceType C5n_Large { get; } = new InstanceType("c5n.large");
        public static InstanceType C5n_XLarge { get; } = new InstanceType("c5n.xlarge");
        public static InstanceType D2_2XLarge { get; } = new InstanceType("d2.2xlarge");
        public static InstanceType D2_4XLarge { get; } = new InstanceType("d2.4xlarge");
        public static InstanceType D2_8XLarge { get; } = new InstanceType("d2.8xlarge");
        public static InstanceType D2_XLarge { get; } = new InstanceType("d2.xlarge");
        public static InstanceType F1_16XLarge { get; } = new InstanceType("f1.16xlarge");
        public static InstanceType F1_2XLarge { get; } = new InstanceType("f1.2xlarge");
        public static InstanceType G2_2XLarge { get; } = new InstanceType("g2.2xlarge");
        public static InstanceType G2_8XLarge { get; } = new InstanceType("g2.8xlarge");
        public static InstanceType G3_16XLarge { get; } = new InstanceType("g3.16xlarge");
        public static InstanceType G3_4XLarge { get; } = new InstanceType("g3.4xlarge");
        public static InstanceType G3_8XLarge { get; } = new InstanceType("g3.8xlarge");
        public static InstanceType G3s_XLarge { get; } = new InstanceType("g3s.xlarge");
        public static InstanceType H1_16XLarge { get; } = new InstanceType("h1.16xlarge");
        public static InstanceType H1_2XLarge { get; } = new InstanceType("h1.2xlarge");
        public static InstanceType H1_4XLarge { get; } = new InstanceType("h1.4xlarge");
        public static InstanceType H1_8XLarge { get; } = new InstanceType("h1.8xlarge");
        public static InstanceType Hs1_8XLarge { get; } = new InstanceType("hs1.8xlarge");
        public static InstanceType I3_16XLarge { get; } = new InstanceType("i3.16xlarge");
        public static InstanceType I3_2XLarge { get; } = new InstanceType("i3.2xlarge");
        public static InstanceType I3_4XLarge { get; } = new InstanceType("i3.4xlarge");
        public static InstanceType I3_8XLarge { get; } = new InstanceType("i3.8xlarge");
        public static InstanceType I3_Large { get; } = new InstanceType("i3.large");
        public static InstanceType I3_XLarge { get; } = new InstanceType("i3.xlarge");
        public static InstanceType I3_Metal { get; } = new InstanceType("i3.metal");
        public static InstanceType M3_2XLarge { get; } = new InstanceType("m3.2xlarge");
        public static InstanceType M3_Large { get; } = new InstanceType("m3.large");
        public static InstanceType M3_Medium { get; } = new InstanceType("m3.medium");
        public static InstanceType M3_XLarge { get; } = new InstanceType("m3.xlarge");
        public static InstanceType M4_10XLarge { get; } = new InstanceType("m4.10xlarge");
        public static InstanceType M4_16XLarge { get; } = new InstanceType("m4.16xlarge");
        public static InstanceType M4_2XLarge { get; } = new InstanceType("m4.2xlarge");
        public static InstanceType M4_4XLarge { get; } = new InstanceType("m4.4xlarge");
        public static InstanceType M4_Large { get; } = new InstanceType("m4.large");
        public static InstanceType M4_XLarge { get; } = new InstanceType("m4.xlarge");
        public static InstanceType M5_Large { get; } = new InstanceType("m5.large");
        public static InstanceType M5_XLarge { get; } = new InstanceType("m5.xlarge");
        public static InstanceType M5_2XLarge { get; } = new InstanceType("m5.2xlarge");
        public static InstanceType M5_4XLarge { get; } = new InstanceType("m5.4xlarge");
        public static InstanceType M5_12XLarge { get; } = new InstanceType("m5.12xlarge");
        public static InstanceType M5_24XLarge { get; } = new InstanceType("m5.24xlarge");
        public static InstanceType M5d_Large { get; } = new InstanceType("m5d.large");
        public static InstanceType M5d_XLarge { get; } = new InstanceType("m5d.xlarge");
        public static InstanceType M5d_2XLarge { get; } = new InstanceType("m5d.2xlarge");
        public static InstanceType M5d_4XLarge { get; } = new InstanceType("m5d.4xlarge");
        public static InstanceType M5d_12XLarge { get; } = new InstanceType("m5d.12xlarge");
        public static InstanceType M5d_24XLarge { get; } = new InstanceType("m5d.24xlarge");
        public static InstanceType M5a_12XLarge { get; } = new InstanceType("m5a.12xlarge");
        public static InstanceType M5a_24XLarge { get; } = new InstanceType("m5a.24xlarge");
        public static InstanceType M5a_2XLarge { get; } = new InstanceType("m5a.2xlarge");
        public static InstanceType M5a_4XLarge { get; } = new InstanceType("m5a.4xlarge");
        public static InstanceType M5a_Large { get; } = new InstanceType("m5a.large");
        public static InstanceType M5a_XLarge { get; } = new InstanceType("m5a.xlarge");
        public static InstanceType P2_16XLarge { get; } = new InstanceType("p2.16xlarge");
        public static InstanceType P2_8XLarge { get; } = new InstanceType("p2.8xlarge");
        public static InstanceType P2_XLarge { get; } = new InstanceType("p2.xlarge");
        public static InstanceType P3_16XLarge { get; } = new InstanceType("p3.16xlarge");
        public static InstanceType P3_2XLarge { get; } = new InstanceType("p3.2xlarge");
        public static InstanceType P3_8XLarge { get; } = new InstanceType("p3.8xlarge");
        public static InstanceType P3dn_24XLarge { get; } = new InstanceType("p3dn.24xlarge");
        public static InstanceType R3_2XLarge { get; } = new InstanceType("r3.2xlarge");
        public static InstanceType R3_4XLarge { get; } = new InstanceType("r3.4xlarge");
        public static InstanceType R3_8XLarge { get; } = new InstanceType("r3.8xlarge");
        public static InstanceType R3_Large { get; } = new InstanceType("r3.large");
        public static InstanceType R3_XLarge { get; } = new InstanceType("r3.xlarge");
        public static InstanceType R4_16XLarge { get; } = new InstanceType("r4.16xlarge");
        public static InstanceType R4_2XLarge { get; } = new InstanceType("r4.2xlarge");
        public static InstanceType R4_4XLarge { get; } = new InstanceType("r4.4xlarge");
        public static InstanceType R4_8XLarge { get; } = new InstanceType("r4.8xlarge");
        public static InstanceType R4_Large { get; } = new InstanceType("r4.large");
        public static InstanceType R4_XLarge { get; } = new InstanceType("r4.xlarge");
        public static InstanceType R5_12XLarge { get; } = new InstanceType("r5.12xlarge");
        public static InstanceType R5_24XLarge { get; } = new InstanceType("r5.24xlarge");
        public static InstanceType R5_2XLarge { get; } = new InstanceType("r5.2xlarge");
        public static InstanceType R5_4XLarge { get; } = new InstanceType("r5.4xlarge");
        public static InstanceType R5_Large { get; } = new InstanceType("r5.large");
        public static InstanceType R5_XLarge { get; } = new InstanceType("r5.xlarge");
        public static InstanceType R5a_12XLarge { get; } = new InstanceType("r5a.12xlarge");
        public static InstanceType R5a_24XLarge { get; } = new InstanceType("r5a.24xlarge");
        public static InstanceType R5a_2XLarge { get; } = new InstanceType("r5a.2xlarge");
        public static InstanceType R5a_4XLarge { get; } = new InstanceType("r5a.4xlarge");
        public static InstanceType R5a_Large { get; } = new InstanceType("r5a.large");
        public static InstanceType R5a_XLarge { get; } = new InstanceType("r5a.xlarge");
        public static InstanceType R5d_12XLarge { get; } = new InstanceType("r5d.12xlarge");
        public static InstanceType R5d_24XLarge { get; } = new InstanceType("r5d.24xlarge");
        public static InstanceType R5d_2XLarge { get; } = new InstanceType("r5d.2xlarge");
        public static InstanceType R5d_4XLarge { get; } = new InstanceType("r5d.4xlarge");
        public static InstanceType R5d_Large { get; } = new InstanceType("r5d.large");
        public static InstanceType R5d_XLarge { get; } = new InstanceType("r5d.xlarge");
        public static InstanceType T2_2XLarge { get; } = new InstanceType("t2.2xlarge");
        public static InstanceType T2_Large { get; } = new InstanceType("t2.large");
        public static InstanceType T2_Medium { get; } = new InstanceType("t2.medium");
        public static InstanceType T2_Micro { get; } = new InstanceType("t2.micro");
        public static InstanceType T2_Nano { get; } = new InstanceType("t2.nano");
        public static InstanceType T2_Small { get; } = new InstanceType("t2.small");
        public static InstanceType T2_XLarge { get; } = new InstanceType("t2.xlarge");
        public static InstanceType T3_2XLarge { get; } = new InstanceType("t3.2xlarge");
        public static InstanceType T3_Large { get; } = new InstanceType("t3.large");
        public static InstanceType T3_Medium { get; } = new InstanceType("t3.medium");
        public static InstanceType T3_Micro { get; } = new InstanceType("t3.micro");
        public static InstanceType T3_Nano { get; } = new InstanceType("t3.nano");
        public static InstanceType T3_Small { get; } = new InstanceType("t3.small");
        public static InstanceType T3_XLarge { get; } = new InstanceType("t3.xlarge");
        public static InstanceType T3a_2XLarge { get; } = new InstanceType("t3a.2xlarge");
        public static InstanceType T3a_Large { get; } = new InstanceType("t3a.large");
        public static InstanceType T3a_Medium { get; } = new InstanceType("t3a.medium");
        public static InstanceType T3a_Micro { get; } = new InstanceType("t3a.micro");
        public static InstanceType T3a_Nano { get; } = new InstanceType("t3a.nano");
        public static InstanceType T3a_Small { get; } = new InstanceType("t3a.small");
        public static InstanceType T3a_XLarge { get; } = new InstanceType("t3a.xlarge");
        public static InstanceType U_12tb1Metal { get; } = new InstanceType("u-12tb1.metal");
        public static InstanceType U_6tb1Metal { get; } = new InstanceType("u-6tb1.metal");
        public static InstanceType U_9tb1Metal { get; } = new InstanceType("u-9tb1.metal");
        public static InstanceType X1_16XLarge { get; } = new InstanceType("x1.16xlarge");
        public static InstanceType X1_32XLarge { get; } = new InstanceType("x1.32xlarge");
        public static InstanceType X1e_16XLarge { get; } = new InstanceType("x1e.16xlarge");
        public static InstanceType X1e_2XLarge { get; } = new InstanceType("x1e.2xlarge");
        public static InstanceType X1e_32XLarge { get; } = new InstanceType("x1e.32xlarge");
        public static InstanceType X1e_4XLarge { get; } = new InstanceType("x1e.4xlarge");
        public static InstanceType X1e_8XLarge { get; } = new InstanceType("x1e.8xlarge");
        public static InstanceType X1e_XLarge { get; } = new InstanceType("x1e.xlarge");
        public static InstanceType Z1d_12XLarge { get; } = new InstanceType("z1d.12xlarge");
        public static InstanceType Z1d_2XLarge { get; } = new InstanceType("z1d.2xlarge");
        public static InstanceType Z1d_3XLarge { get; } = new InstanceType("z1d.3xlarge");
        public static InstanceType Z1d_6XLarge { get; } = new InstanceType("z1d.6xlarge");
        public static InstanceType Z1d_Large { get; } = new InstanceType("z1d.large");
        public static InstanceType Z1d_XLarge { get; } = new InstanceType("z1d.xlarge");

        public static bool operator ==(InstanceType left, InstanceType right) => left.Equals(right);
        public static bool operator !=(InstanceType left, InstanceType right) => !left.Equals(right);

        public static explicit operator string(InstanceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InstanceType other && Equals(other);
        public bool Equals(InstanceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The strategy of the placement group determines how the instances are organized within the group.
    /// See https://docs.aws.amazon.com/cli/latest/reference/ec2/create-placement-group.html
    /// </summary>
    [EnumType]
    public readonly struct PlacementStrategy : IEquatable<PlacementStrategy>
    {
        private readonly string _value;

        private PlacementStrategy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// A `spread` placement group places instances on distinct hardware.
        /// </summary>
        public static PlacementStrategy Spread { get; } = new PlacementStrategy("spread");
        /// <summary>
        /// A `cluster` placement group is a logical grouping of instances within a single
        /// Availability Zone that benefit from low network latency, high network throughput.
        /// </summary>
        public static PlacementStrategy Cluster { get; } = new PlacementStrategy("cluster");

        public static bool operator ==(PlacementStrategy left, PlacementStrategy right) => left.Equals(right);
        public static bool operator !=(PlacementStrategy left, PlacementStrategy right) => !left.Equals(right);

        public static explicit operator string(PlacementStrategy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PlacementStrategy other && Equals(other);
        public bool Equals(PlacementStrategy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ProtocolType : IEquatable<ProtocolType>
    {
        private readonly string _value;

        private ProtocolType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ProtocolType All { get; } = new ProtocolType("all");
        public static ProtocolType TCP { get; } = new ProtocolType("tcp");
        public static ProtocolType UDP { get; } = new ProtocolType("udp");
        public static ProtocolType ICMP { get; } = new ProtocolType("icmp");

        public static bool operator ==(ProtocolType left, ProtocolType right) => left.Equals(right);
        public static bool operator !=(ProtocolType left, ProtocolType right) => !left.Equals(right);

        public static explicit operator string(ProtocolType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProtocolType other && Equals(other);
        public bool Equals(ProtocolType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct Tenancy : IEquatable<Tenancy>
    {
        private readonly string _value;

        private Tenancy(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Tenancy Default { get; } = new Tenancy("default");
        public static Tenancy Dedicated { get; } = new Tenancy("dedicated");

        public static bool operator ==(Tenancy left, Tenancy right) => left.Equals(right);
        public static bool operator !=(Tenancy left, Tenancy right) => !left.Equals(right);

        public static explicit operator string(Tenancy value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Tenancy other && Equals(other);
        public bool Equals(Tenancy other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
